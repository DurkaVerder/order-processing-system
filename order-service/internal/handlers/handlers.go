package handlers

import (
	"order-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	HandlerAddOrder(c *gin.Context)
	HandlerGetOrder(c *gin.Context)
	HandlerGetAllOrders(c *gin.Context)
	HandlerDeleteOrder(c *gin.Context)
	HandlerChangeStatus(c *gin.Context)
}

type HandlersManager struct {
	service service.Service
}

func NewHandlersManager(service service.Service) *HandlersManager {
	return &HandlersManager{
		service: service,
	}
}
