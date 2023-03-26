package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lukamindo/pet-reminder/api/handler"
	"github.com/lukamindo/pet-reminder/helper/conn"
)

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

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Router
	handler.New(e)

	// Listen and Serve
	e.Logger.Fatal(e.Start(":1323"))
}
