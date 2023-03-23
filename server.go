package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/lukamindo/pet-reminder/api/handler"
	"github.com/lukamindo/pet-reminder/helper/conn"
	"github.com/lukamindo/pet-reminder/helper/watcher"
)

func main() {
	restart := watcher.GetNotifier()
	go func() {
		<-restart
		log.Printf(" I will restart shortly ...\n")
	}()
	watcher.StartWatcher()

	fmt.Println(conn.New())
	e := echo.New()
	handler.New(e)

	e.Logger.Fatal(e.Start(":1323"))
}
