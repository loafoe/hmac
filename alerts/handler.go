package alerts

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handler returns an Echo handler
func Handler(token string, storer Storer) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := c.Param("token")
		if t != token {
			return c.JSON(http.StatusUnauthorized, nil)
		}
		p := new(Payload)
		if err := c.Bind(p); err != nil {
			return err
		}
		err := storer.Store(*p)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, p)
	}
}
