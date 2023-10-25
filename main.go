package main

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lukamindo/pet-reminder/api/handler"
	"github.com/lukamindo/pet-reminder/helper/conn"
	"github.com/lukamindo/pet-reminder/helper/server"
)

type errorResponse struct {
	Message string `json:"message"`
}

func main() {
	// restart := watcher.GetNotifier()
	// go func() {
	// 	<-restart
	// 	log.Printf(" I will restart shortly ...\n")
	// }()
	// watcher.StartWatcher()

	// Connect to DB
	conn.New()

	// Initialize Echo
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "I am we-api and i am good!")
	})

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.HTTPErrorHandler = ErrorHandler

	// Router
	handler.New(e)

	// Listen and Serve
	e.Logger.Fatal(e.Start(":1323"))
}

func ErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}
	e := errorResponse{
		Message: err.Error(),
	}
	if customErr, ok := err.(server.Err); ok {
		c.Response().WriteHeader(customErr.StatusCode)
	}
	c.Response().WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(c.Response().Writer).Encode(e)
}
