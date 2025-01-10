package main

import (
	"APIGateway/config"
	"APIGateway/internal/handlers"
	"APIGateway/internal/requester"
	"APIGateway/internal/server"
)

func main() {
	config := config.LoadConfig()

	requester := requester.NewRequestManager()

	handlers := handlers.NewHandlersManager(requester, config)

	server := server.NewServer(handlers, config)

	server.Start()
}
