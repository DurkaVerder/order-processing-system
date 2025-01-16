package main

import (
	"log"
	"order-service/config"
	"order-service/internal/handlers"
	"order-service/internal/kafka/producer"
	"order-service/internal/repository/postgres"
	"order-service/internal/server"
	"order-service/internal/service"
	"os"
)

func main() {
	config := config.LoadConfig()

	postgres := postgres.NewPostgres()

	producer, err := producer.NewProducerManager([]string{os.Getenv("KAFKA_BROKERS")})
	if err != nil {
		log.Fatalf("Error while creating producer: %v", err)
	}

	service := service.NewServiceManager(postgres, producer)

	handlers := handlers.NewHandlersManager(service)

	server := server.NewServer(handlers, config)

	server.Start()

}
