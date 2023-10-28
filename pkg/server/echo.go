package server

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/lukamindo/pet-reminder/app/constant"
)

// TODO: amas rame movufiqro
type ContextKey string

const (
	ContextID ContextKey = "ID"
)

func Echo() *echo.Echo {
	e := echo.New()

	// hide Echo banner
	e.HideBanner = true

	// define log level based on env
	e.Logger.SetLevel(logLevel())

	// custom error handler
	e.HTTPErrorHandler = ErrorHandler

	e.Use(middleware.Recover())

	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		RequestIDHandler: func(c echo.Context, id string) {
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), ContextID, id))) // c.Value(key.ContextID).(string)
		},
	}))

	// TODO: logger unda gavaketo chemi
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "lvl=REQ id=${id} message='\033[31m${error}\033[00m' \033[36m${method} • ${uri}\033[00m → status=${status} ip=${remote_ip} latency=${latency_human} \033[31m${myerr}\033[00m${body} ${callstack}\n",
		Output: e.Logger.Output(), // Default value os.Stdout, e.Logger use gommon
	}))
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
