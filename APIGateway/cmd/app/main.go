package main

import (
	"APIGateway/config"
	"APIGateway/internal/handlers"
	"APIGateway/internal/requester"
	"APIGateway/internal/server"
	"net/http"
	"time"
)

func main() {
	config := config.LoadConfig()

	requester := requester.NewRequestManager(&http.Client{Timeout: 10 * time.Second})

	handlers := handlers.NewHandlersManager(requester, config)

	server := server.NewServer(handlers, config)

	server.Start()
}
