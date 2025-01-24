package producer

import (
	"authentication-service/internal/kafka"
	"encoding/json"
	"log"
	"time"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	"github.com/IBM/sarama"
)

// Producer is an interface for Kafka producer
type Producer interface {
	SendMessage(topic string, message common.DataForNotify) error
}

// ProducerManager is a Kafka producer
type ProducerManager struct {
	producer sarama.SyncProducer
	config   *sarama.Config
}

// NewProducerManager creates a new Kafka producer
func NewProducerManager(brokers []string) (*ProducerManager, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	var err error
	var producer sarama.SyncProducer
	for i := 0; i < kafka.MaxRetry; i++ {
		producer, err = sarama.NewSyncProducer(brokers, config)
		if err == nil {
			log.Println("Producer created")
			return &ProducerManager{
				producer: producer,
				config:   config,
			}, nil
		}
		log.Printf("Failed to create producer: %s, retrying...", err)
		time.Sleep(time.Second * 2)
	}
	log.Printf("Failed to create producer: %s", err)

	return nil, err
}

// SendMessage sends a message to a Kafka topic
func (p *ProducerManager) SendMessage(topic string, message common.DataForNotify) error {
	data, err := json.Marshal(message)
	if err != nil {
		log.Printf("Failed to marshal message: %s", err)
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %s", err)
		return err
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}
