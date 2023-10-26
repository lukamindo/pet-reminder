package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lukamindo/pet-reminder/api/handler"
	"github.com/lukamindo/pet-reminder/app/constant"
	"github.com/lukamindo/pet-reminder/helper/conn"
	"github.com/lukamindo/pet-reminder/helper/server"
)

type errorResponse struct {
	Message string `json:"message"`
}

func main() {
	// Connect to DB
	conn.New()

	// Initialize Echo //TODO: echo ragaceebi calke unda gavitano
	e := echo.New()

	// health endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "i am pet reminder api and i am good!")
	})

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.HTTPErrorHandler = ErrorHandler

	// Router
	handler.New(e)

	// Listen and Serve
	go server.Start(e.Start, os.Getenv(constant.Environment) == "dev")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		// deferCB()
		cancel()
	}()
	if err := e.Shutdown(ctx); err != nil {
		log.Panic(err)
	}
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
