package main

import (
	"echo/config"
	"echo/internal/server"
	"log"
)

func main() {
	server := server.NewServer(config.ServerAddr)

	if err := server.Listen(); err != nil {
		log.Fatalf("server listen with err: %v", err)
	}

	server.Run()
}
