package handlers

import (
	"APIGateway/config"

	"github.com/gin-gonic/gin"
)

const (
	StartURL = "http://localhost"
)

// Handlers is an interface that defines the methods that the handlers must implement
type Handlers interface {
	HandlerLogin(c *gin.Context)
	HandlerRegister(c *gin.Context)
	HandlerLogout(c *gin.Context)
	HandlerCreateOrder(c *gin.Context)
	HandlerGetOrders(c *gin.Context)
	HandlerGetOrder(c *gin.Context)
	HandlerDeleteOrder(c *gin.Context)
	HandlerStatusOrder(c *gin.Context)
	HandlerHistoryOrder(c *gin.Context)
}

type HandlersManager struct {
	cfg config.Config
}

func NewHandlersManager(cfg config.Config) Handlers {
	return &HandlersManager{
		cfg: cfg,
	}
}
