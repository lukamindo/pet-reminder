package server

import (
	"log"
	"os"

	"github.com/lukamindo/pet-reminder/app/constant"
	"github.com/lukamindo/pet-reminder/helper/watcher"
)

type Starter func(string) error

func Start(start Starter, shouldWatch bool) {
	if shouldWatch {
		restart := watcher.GetNotifier()
		go func() {
			<-restart
			log.Printf("I will restart shortly...\n")
		}()
		watcher.StartWatcher()

	}
	if err := start(os.Getenv(constant.ServerPort)); err != nil {
		log.Fatalf("shutting down the server, error: %s\n", err)
	}
}
