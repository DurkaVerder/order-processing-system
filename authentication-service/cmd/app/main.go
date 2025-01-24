package main

import (
	"authentication-service/config"
	"authentication-service/internal/handlers"
	"authentication-service/internal/kafka/producer"
	"authentication-service/internal/repository/postgres"
	"authentication-service/internal/repository/redis"
	"authentication-service/internal/server"
	"authentication-service/internal/service"
	"log"
	"os"
)

func main() {
	cfg := config.LoadConfig()

	postgres := postgres.NewPostgres()

	redis := redis.NewRedis()

	producer, err := producer.NewProducerManager([]string{os.Getenv("KAFKA_BROKERS")})
	if err != nil {
		log.Fatalf("Error while creating producer: %v", err)
	}

	service := service.NewServiceManager(postgres, redis, producer)

	handlers := handlers.NewHandlersManager(service)

	server := server.NewServer(handlers, cfg)

	server.Start()
}
