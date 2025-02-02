package service

import (
	"errors"
	"log"
	"status-service/internal/kafka/producer"
	"status-service/internal/repository"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
)

type Service interface {
	ChangeStatus(orderId int, status string) error
}

type ServiceManager struct {
	db       repository.DataBase
	producer producer.Producer
}

func NewServiceManager(db repository.DataBase, producer producer.Producer) *ServiceManager {
	return &ServiceManager{
		db:       db,
		producer: producer,
	}
}

func (s *ServiceManager) ChangeStatus(orderId int, status string) error {
	if orderId <= 0 {
		return errors.New("invalid orderId")
	}

	if status == "" {
		return errors.New("invalid status")
	}

	err := s.db.UpdateStatus(orderId, status)
	if err != nil {
		return err
	}

	if err := s.db.CreateRecordStatus(orderId, status); err != nil {
		return err
	}

	go func() {
		notify := s.createMessage(orderId, status)
		if err := s.producer.SendMessage("notification", notify); err != nil {
			log.Printf("Error while sending message: %v", err)
		}
	}()
	return nil
}

func (s *ServiceManager) createMessage(orderId int, status string) common.DataForNotify {
	notify := common.DataForNotify{
		Event:   "update_status",
		OrderId: orderId,
		Status:  status,
	}
	return notify
}
