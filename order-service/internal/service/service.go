package service

import (
	"fmt"
	"log"
	"order-service/internal/kafka"
	"order-service/internal/kafka/producer"
	"order-service/internal/repository"
	"strconv"
	"time"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
)

type Service interface {
	AddOrder(order common.Order) error
	GetOrder(id string) (common.Order, error)
	GetAllOrders(userId string) ([]common.Order, error)
	DeleteOrder(id string) error
	ChangeStatus(orderId, status string) error
}

type ServiceManager struct {
	db       repository.DateBase
	producer producer.Producer
}

func NewServiceManager(db repository.DateBase, producer producer.Producer) *ServiceManager {
	return &ServiceManager{
		db:       db,
		producer: producer,
	}
}

func (s *ServiceManager) AddOrder(order common.Order) error {
	order = s.initOrder(time.Now(), order)

	err := s.db.AddOrder(order)
	if err != nil {
		return err
	}

	go func() {
		msg := s.createMessage(order)

		if err := s.producer.SendMessage(kafka.NotificationTopic, msg); err != nil {
			log.Printf("Error while sending message: %v", err)
			return
		}
		log.Printf("Message sent: %v", msg)
	}()

	return nil
}

func (s *ServiceManager) initOrder(time time.Time, order common.Order) common.Order {
	order.CreatedAt = time
	order.UpdateAt = time
	order.Status = "created"

	return order
}

func (s *ServiceManager) createMessage(order common.Order) common.DataForNotify {
	return common.DataForNotify{
		Event:   "order_created",
		OrderId: order.Id,
		Status:  "created",
	}
}

func (s *ServiceManager) GetOrder(id string) (common.Order, error) {
	return s.db.GetOrder(id)
}

func (s *ServiceManager) GetAllOrders(userId string) ([]common.Order, error) {
	return s.db.GetAllOrders(userId)
}

func (s *ServiceManager) DeleteOrder(id string) error {
	return s.db.DeleteOrder(id)
}

func (s *ServiceManager) ChangeStatus(orderId, status string) error {
	id, err := strconv.Atoi(orderId)
	if err != nil {
		fmt.Printf("Error while parsing orderId: %v", err)
		return err
	}

	go func() {
		msg := common.StatusOrder{
			OrderId:   id,
			Status:    status,
			CreatedAt: time.Now(),
		}

		if err := s.producer.SendMessage(kafka.StatusTopic, msg); err != nil {
			log.Printf("Error while sending message: %v", err)
			return
		}
		log.Printf("Message sent: %v", msg)
	}()

	return nil
}
