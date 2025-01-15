package producer

import (
	"authentication-service/internal/kafka"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"
)

type Producer interface {
	SendMessage(topic string, message any) error
}

type ProducerManager struct {
	producer sarama.SyncProducer
	config   *sarama.Config
}

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
	}
	log.Printf("Failed to create producer: %s", err)

	return nil, err
}

func (p *ProducerManager) SendMessage(topic string, message any) error {
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
