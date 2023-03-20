package server

import (
	"net/http"

	"github.com/labstack/echo"
)

// Success wraps payoad in data field and responds with corresponding json
func Success(c echo.Context, payload interface{}) error {
	return c.JSON(http.StatusOK, struct {
		Data interface{} `json:"data"`
	}{Data: payload})
}

// JSON wraps payoad in data field and responds with corresponding json under status code specified
func JSON(c echo.Context, statusCode int, payload interface{}) error {
	return c.JSON(statusCode, struct {
		Data interface{} `json:"data"`
	}{Data: payload})
}
