package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

func ErrBadRequest(err error) Err {
	return bootstrap(err, ErrorBadRequest)
}

func ErrUnauthorized(err error) Err {
	return bootstrap(err, ErrorUnauthorized)
}

func ErrForbidden(err error) Err {
	return bootstrap(err, ErrorForbidden)
}

func ErrInternalDB(err error) Err {
	return bootstrap(err, ErrorInternalDB)
}

func ErrInternalDomain(err error) Err {
	return bootstrap(err, ErrorInternalDomain)
}

func bootstrap(err error, costumeErr Err) Err {
	costumeErr.Message = err.Error()
	return costumeErr
}
