package main

import (
	"log"
	"os"
	"status-service/internal/kafka"
	"status-service/internal/kafka/consumer"
	"status-service/internal/kafka/producer"
	"status-service/internal/repository/postgres"
	"status-service/internal/service"
)

func main() {
	postgres := postgres.NewPostgres()

	producer, err := producer.NewProducerManager([]string{os.Getenv("KAFKA_BROKERS")})
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}

	service := service.NewServiceManager(postgres, producer)

	consumer := consumer.NewConsumerManager([]string{os.Getenv("KAFKA_BROKERS")}, service)
	consumer.Subscribe(kafka.StatusTopic)
	consumer.Start()
}
