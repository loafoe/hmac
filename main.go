package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/philips-software/hsdp-metrics-alert-collector/alerts"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	storer, err := alerts.NewPGStorer()
	if err != nil {
		log.Fatal(err)
	}

	e.POST("/alerts", alerts.Handler(storer))

	e.Logger.Fatal(e.Start(":1323"))

}
