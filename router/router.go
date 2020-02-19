package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/philips-labs/hmac/alerts"
)

// New returns a new router
func New(storer alerts.Storer) *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/alerts", alerts.Handler(storer))

	return e
}
