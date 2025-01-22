package main

import (
	"context"
	"log"
	"notification-service/internal/kafka"
	"notification-service/internal/kafka/consumer"
	"notification-service/internal/repository/postgres"
	"os/signal"
	"syscall"

	"notification-service/internal/service"
	"os"
)

func main() {
	postgres := postgres.NewPostgres()

	service := service.NewServiceManager(postgres)

	consumer := consumer.NewConsumerManager([]string{os.Getenv("KAFKA_BROKER")}, service)
	consumer.Subscribe(kafka.NotificationTopic)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	consumer.Start(ctx)

	service.StarProcessingAndSendingMsg(5, ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-signalChan
	log.Printf("Received signal: %v. Shutting down...\n", sig)
	service.StopProcessingAndSendingMsg()
}
