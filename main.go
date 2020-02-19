package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/markbates/pkger"
	"github.com/philips-labs/hmac/alerts"
	"github.com/philips-labs/hmac/router"
)

//go:generate pkger -include /migrations

func main() {
	pkger.Walk("/migrations", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout,
			"%s \t %d \t %s \t %s \t\n",
			info.Name(),
			info.Size(),
			info.Mode(),
			info.ModTime().Format(time.RFC3339),
		)

		return nil
	})

	storer, err := alerts.NewPGStorer()
	if err != nil {
		log.Fatal(err)
	}
	e := router.New(storer)
	e.Logger.Fatal(e.Start(":8080"))
}
