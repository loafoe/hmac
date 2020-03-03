package alerts

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type errorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func jsonError(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, &errorResponse{
		StatusCode: statusCode,
		Message:    message,
	})
}

// AddHandler returns an add handler
func AddHandler(token string, storer Storer) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := c.Param("token")
		if t != token {
			return jsonError(c, http.StatusUnauthorized, "invalid token")
		}
		p := new(Payload)
		if err := c.Bind(p); err != nil {
			return jsonError(c, http.StatusBadRequest, "parsing error: "+err.Error())
		}
		if p.AlertName == "" {
			return jsonError(c, http.StatusBadRequest, "bad request: missing alertName")
		}
		// Setup init payload
		if len(p.Alerts) == 0 {
			p.Alerts = []Alert{
				Alert{
					Status: "resolved",
					Labels: map[string]string{
						"alertname":    p.AlertName,
						"status":       "resolved",
						"region":       "unknown",
						"organization": "unknown",
						"space":        "unknown",
					},
					Annotations: map[string]string{
						"description": "Initialization record",
					},
				},
			}
			p.GroupLabels = map[string]string{
				"alertname": p.AlertName,
			}
		}
		err := storer.Store(*p)
		if err != nil {
			return jsonError(c, http.StatusInternalServerError, "storing error: "+err.Error())
		}
		return c.JSON(http.StatusOK, p)
	}
}

// DeleteHandler returns a delete handler
func DeleteHandler(token string, storer Storer) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := c.Param("token")
		if t != token {
			return jsonError(c, http.StatusUnauthorized, "invalid token")
		}
		p := new(Payload)
		if err := c.Bind(p); err != nil {
			return jsonError(c, http.StatusBadRequest, "parsing error: "+err.Error())
		}
		if p.AlertName == "" {
			return jsonError(c, http.StatusBadRequest, "bad request: missing alertName")
		}
		// Delete payloads
		err := storer.Remove(*p)
		if err != nil {
			return jsonError(c, http.StatusInternalServerError, "delete error: "+err.Error())
		}
		return c.JSON(http.StatusNoContent, "")
	}
}

// StoreHandler returns an Echo handler for storing payloads
func StoreHandler(token string, storer Storer) echo.HandlerFunc {
	return func(c echo.Context) error {
		t := c.Param("token")
		if t != token {
			return jsonError(c, http.StatusUnauthorized, "invalid token")
		}
		p := new(Payload)
		if err := c.Bind(p); err != nil {
			return jsonError(c, http.StatusBadRequest, "parsing error: "+err.Error())
		}
		if p.Status == "" {
			return jsonError(c, http.StatusBadRequest, "bad request: missing status")
		}
		if p.AlertName == "" && len(p.Alerts) == 0 && len(p.GroupLabels) == 0 {
			return jsonError(c, http.StatusBadRequest, "bad request: missing alertName")
		}
		err := storer.Store(*p)
		if err != nil {
			return jsonError(c, http.StatusInternalServerError, "storing error: "+err.Error())
		}
		return c.JSON(http.StatusOK, p)
	}
}
