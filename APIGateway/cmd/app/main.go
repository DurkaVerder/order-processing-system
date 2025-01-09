package main

import (
	"APIGateway/config"
	"APIGateway/internal/handlers"
	"APIGateway/internal/server"
)

func main() {
	config := config.LoadConfig()

	handlers := handlers.NewHandlersManager(config)

	server := server.NewServer(handlers)

	server.Start(config.Gateway.Server.Port)
}
