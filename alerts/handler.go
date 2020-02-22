package alerts

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	statusCode int    `json:"statusCode"`
	message    string `json:"message"`
}

func jsonError(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, &errorResponse{
		statusCode: statusCode,
		message:    message,
	})
}

// Handler returns an Echo handler
func Handler(token string, storer Storer) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := c.Param("token")
		if t != token {
			return jsonError(c, http.StatusUnauthorized, "invalid token")
		}
		p := new(Payload)
		if err := c.Bind(p); err != nil {
			return jsonError(c, http.StatusBadRequest, "parsing error")
		}
		if p.Status == "" {
			return jsonError(c, http.StatusBadRequest, "bad request")
		}
		err := storer.Store(*p)
		if err != nil {
			return jsonError(c, http.StatusInternalServerError, "storing error")
		}
		return c.JSON(http.StatusOK, p)
	}
}
