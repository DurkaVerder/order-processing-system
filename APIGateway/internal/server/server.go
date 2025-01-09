package server

import (
	"APIGateway/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handlers handlers.Handlers
}

func NewServer(handlers handlers.Handlers) *Server {
	return &Server{
		handlers: handlers,
	}
}

func (s *Server) Start(port string) {
	r := gin.Default()

	r.Run(port)
}
