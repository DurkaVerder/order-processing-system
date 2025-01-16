package server

import (
	"order-service/config"
	"order-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handlers handlers.Handlers
	config   config.Config
}

func NewServer(handlers handlers.Handlers, config config.Config) *Server {
	return &Server{
		handlers: handlers,
		config:   config,
	}
}

func (s *Server) Start() {
	r := gin.Default()

	order := r.Group(s.config.Order.Route.Base)
	{
		order.POST(s.config.Order.Route.Endpoints["create_order"], s.handlers.HandlerAddOrder)
		order.GET(s.config.Order.Route.Endpoints["get_order"], s.handlers.HandlerGetOrder)
		order.GET(s.config.Order.Route.Endpoints["get_orders"], s.handlers.HandlerGetAllOrders)
		order.DELETE(s.config.Order.Route.Endpoints["delete_order"], s.handlers.HandlerDeleteOrder)
	}

	r.Run(s.config.Order.Server.Port)
}
