package handlers

import (
	"authentication-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Logout(c *gin.Context)
	ValidateToken(c *gin.Context)
}

type HandlersManager struct {
	service service.Service
}

func NewHandlersManager(service service.Service) *HandlersManager {
	return &HandlersManager{
		service: service,
	}
}
