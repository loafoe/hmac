package alerts

import (
	"encoding/json"
	"log"
	"time"

	"github.com/cloudfoundry-community/gautocloud"
	_ "github.com/cloudfoundry-community/gautocloud/connectors/databases/client/postgresql" // PostgreSQL connector
	"github.com/cloudfoundry-community/gautocloud/connectors/databases/dbtype"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/types"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/philips-labs/hmac/migrations"
)

// PGPayload is an entry in the alerts table that wraps Payload
type PGPayload struct {
	ID        int64          `db:"id"`
	CreatedAt time.Time      `db:"created_at"`
	Payload   types.JSONText `db:"payload"`
}

// PGStorer is an PostgreSQL implementation of the Storer interface
type PGStorer struct {
	svc *dbtype.PostgresqlDB
	db  *sqlx.DB
}

// Init initializes the database
func (p *PGStorer) Init() error {
	// wrap assets into Resource
	assetSource := bindata.Resource(migrations.AssetNames(),
		func(name string) ([]byte, error) {
			return migrations.Asset(name)
		})
	driver, err := postgres.WithInstance(p.db.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	d, err := bindata.WithInstance(assetSource)
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance("go-bindata", d, "postgres://database", driver)
	return m.Up()
}

// Remove removes all instances of alertName
func (p *PGStorer) Remove(payload Payload) error {
	db := p.db

	tx := db.MustBegin()

	tx.MustExec("DELETE FROM alerts WHERE (payload->>'alertName' = $1 OR payload->'alerts'->0->'labels'->>'alertname' = $1)", payload.AlertName)
	return tx.Commit()
}

// Store stores a payload in the psql DB
func (p *PGStorer) Store(payload Payload) error {
	db := p.db

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	tx := db.MustBegin()

	tx.MustExec("INSERT INTO alerts (payload) VALUES ($1)", payloadJSON)
	return tx.Commit()
}

// NewPGStorer retruns a psql storer
func NewPGStorer() (*PGStorer, error) {
	storer := &PGStorer{}

	err := gautocloud.Inject(&storer.svc)

	if err != nil {
		return nil, err
	}

	storer.db = sqlx.NewDb(storer.svc.DB, "postgres")
	err = storer.db.Ping()
	if err != nil {
		return nil, err
	}

	return storer, nil
}
