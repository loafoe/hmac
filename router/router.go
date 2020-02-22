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
func New(config Config) *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/webhook/alerts/:token", alerts.Handler(config.Token, config.Storer))

	return e
}
