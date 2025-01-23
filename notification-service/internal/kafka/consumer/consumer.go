package consumer

import (
	"context"
	"encoding/json"
	"log"
	"notification-service/internal/kafka"
	"notification-service/internal/service"
	"time"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	"github.com/IBM/sarama"
)

type Consumer interface {
	Subscribe(topic string) error
	Start()
}

type ConsumerManager struct {
	consumer         sarama.Consumer
	config           sarama.Config
	consumePartition sarama.PartitionConsumer
	service          service.Service
}

func NewConsumerManager(brokers []string, service service.Service) *ConsumerManager {
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	for i := 0; i < kafka.MaxRetry; i++ {
		consumer, err := sarama.NewConsumer(brokers, config)
		if err == nil {
			return &ConsumerManager{
				consumer: consumer,
				config:   *config,
				service:  service,
			}
		}
		log.Printf("Failed to create consumer: %s, retrying...", err)
		time.Sleep(time.Second * 2)
	}

	log.Fatalln("Failed to create consumer")
	return nil
}

func (c *ConsumerManager) Subscribe(topic string) error {
	var err error
	c.consumePartition, err = c.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Printf("Failed to consume partition: %s", err)
		return err
	}
	return nil
}

func (c *ConsumerManager) Start(ctx context.Context) {
	for {
		select {
		case msg := <-c.consumePartition.Messages():
			log.Printf("Received message: %s", msg.Value)

			notify := common.DataForNotify{}
			if err := json.Unmarshal(msg.Value, &notify); err != nil {
				log.Printf("Failed to unmarshal message: %s", err)
				return
			}
			c.service.AddDataForNotifyInChan(notify)

		case err := <-c.consumePartition.Errors():
			log.Printf("Error: %s", err)

		case <-ctx.Done():
			log.Println("Consumer stopping: context cancelled")
			return
		}

	}
}
