package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	ticker := time.Tick(time.Second * 5)

	for {
		select {
		case <- ticker:
			log.Println("Prints Event 5 seconds ")
		case <-interrupt:
			log.Println("Shutting down.")
			return
		}
	}
}
