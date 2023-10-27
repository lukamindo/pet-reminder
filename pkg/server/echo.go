package server

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/lukamindo/pet-reminder/app/constant"
)

func Echo() *echo.Echo {
	e := echo.New()

	// hide Echo banner
	e.HideBanner = true

	// define log level based on env
	e.Logger.SetLevel(logLevel())

	e.HTTPErrorHandler = ErrorHandler

	e.Use(middleware.Recover())

	//TODO: test all middlewares
	// e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	// e.Validator = validator.New()
	return e
}

func logLevel() log.Lvl {
	if os.Getenv(constant.Environment) == "dev" {
		return log.DEBUG
	}
	return log.WARN
}

type errorResponse struct {
	Message string `json:"message"`
}

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}
	errResponse := errorResponse{
		Message: err.Error(),
	}
	customErr, ok := err.(Err)
	if !ok {
		c.Response().WriteHeader(http.StatusInternalServerError)
	}
	c.Response().WriteHeader(customErr.StatusCode)
	_ = json.NewEncoder(c.Response().Writer).Encode(errResponse)
}
