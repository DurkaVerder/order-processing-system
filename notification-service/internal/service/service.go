package service

import (
	"notification-service/internal/repository"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
)

type Service interface {
	CreateNotification(notify common.DataForNotify) (common.Notification, error)
	SendNotification(notify common.Notification) error
}

type ServiceManager struct {
	db repository.DataBase
}

func NewServiceManager(db repository.DataBase) *ServiceManager {
	return &ServiceManager{
		db: db,
	}
}

func (s *ServiceManager) CreateNotification(notify common.DataForNotify) (common.Notification, error) {
	notification := common.Notification{}
	switch notify.Event {
	case "order_created":
		notification.Subject = "Order created"
		notification.Body = "Order has been created"

	case "order_updated":
		notification.Subject = "Order updated"
		notification.Body = "Order has been updated"
	}
	email, err := s.db.GetUserEmailByOrderId(notify.OrderId)
	if err != nil {
		return common.Notification{}, err
	}
	notification.To = email

	return notification, nil
}

func (s *ServiceManager) SendNotification(notify common.Notification) error {

	return nil
}
