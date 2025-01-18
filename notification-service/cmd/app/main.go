package main

import (
	"notification-service/internal/kafka"
	"notification-service/internal/kafka/consumer"
	"notification-service/internal/repository/postgres"

	"notification-service/internal/service"
	"os"
)

func main() {
	postgres := postgres.NewPostgres()

	service := service.NewServiceManager(postgres)

	consumer := consumer.NewConsumerManager([]string{os.Getenv("KAFKA_BROKER")}, service)
	consumer.Subscribe(kafka.NotificationTopic)
	consumer.Start()
}
