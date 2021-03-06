package main

import (
	"echo/config"
	"echo/internal/client"
	"log"
)

func main() {
	client := client.NewClient(config.ServerAddr)

	if err := client.Connect(); err != nil {
		log.Fatalf("client connect with err: %v", err)
	}

	client.Run()
}
