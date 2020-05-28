package main

import (
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			log.Println("[twitter-craler] run...")
		}
	}
}
