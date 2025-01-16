package repository

import common "github.com/DurkaVerder/common-for-order-processing-system/models"

type DateBase interface {
	AddOrder(order common.Order) error
	GetOrder(id string) (common.Order, error)
	GetAllOrders(userId string) ([]common.Order, error)
	DeleteOrder(id string) error
	GetUserEmail(userId string) (string, error)
}
