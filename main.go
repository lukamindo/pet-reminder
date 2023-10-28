package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/lukamindo/pet-reminder/api/handler"
	"github.com/lukamindo/pet-reminder/app/constant"
	"github.com/lukamindo/pet-reminder/pkg/conn"
	"github.com/lukamindo/pet-reminder/pkg/server"
)

//TODO:
/*

1. echo gavitano calke +++
2. logger unda davamato
3. validator მივაბა echos


*/
func init() {
	// Connect to DB
	conn.New()

}

func main() {
	// Initialize Echo
	e := server.Echo()

	// health endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "i am pet reminder api and i am good!")
	})

	// Router
	handler.New(e)

	// Start Server
	go server.Start(e.Start, os.Getenv(constant.Environment) == "dev")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		log.Panic(err)
	}
}
