package handlers

import (
	"APIGateway/config"
	"APIGateway/internal/requester"

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
	HandlerHistoryOrder(c *gin.Context)
}

// HandlersManager is a struct that implements Handlers interface
type HandlersManager struct {
	requester requester.Requester
	cfg       config.Config
}

// NewHandlersManager is a constructor for HandlersManager
func NewHandlersManager(requester requester.Requester, cfg config.Config) *HandlersManager {
	return &HandlersManager{
		requester: requester,
		cfg:       cfg,
	}
}
