package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/philips-labs/hmac/alerts"
)

type Config struct {
	Storer alerts.Storer
	Token  string
}

// New returns a new router
func New(config Config) (*echo.Echo, error) {
	// Init storer
	err := config.Storer.Init()
	if err != nil {
		return nil, err
	}

	// create a new echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/webhook/alerts/:token", alerts.StoreHandler(config.Token, config.Storer))
	e.PUT("/webhook/alerts/:token", alerts.AddHandler(config.Token, config.Storer))
	e.DELETE("/webhook/alerts/:token", alerts.DeleteHandler(config.Token, config.Storer))

	return e, nil
}
