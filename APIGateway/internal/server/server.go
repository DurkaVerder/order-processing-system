package server

import (
	"APIGateway/config"
	"APIGateway/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	handlers handlers.Handlers
	cfg      config.Config
}

func NewServer(handlers handlers.Handlers, cfg config.Config) *Server {
	return &Server{
		handlers: handlers,
		cfg:      cfg,
	}
}

func (s *Server) Start() {
	r := gin.Default()

	protected := r.Group(s.cfg.Gateway.Route.Base)
	protected.Use(s.authMiddleware)
	{
		protected.GET(s.cfg.Gateway.Route.Endpoints["logout"], s.handlers.HandlerLogout)
		protected.POST(s.cfg.Gateway.Route.Endpoints["create_order"], s.handlers.HandlerCreateOrder)
		protected.GET(s.cfg.Gateway.Route.Endpoints["get_orders"], s.handlers.HandlerGetOrders)
		protected.GET(s.cfg.Gateway.Route.Endpoints["get_order"], s.handlers.HandlerGetOrder)
		protected.DELETE(s.cfg.Gateway.Route.Endpoints["delete_order"], s.handlers.HandlerDeleteOrder)
		protected.GET(s.cfg.Gateway.Route.Endpoints["status_order"], s.handlers.HandlerStatusOrder)
		protected.GET(s.cfg.Gateway.Route.Endpoints["history_order"], s.handlers.HandlerHistoryOrder)
	}
	auth := r.Group(s.cfg.Gateway.Route.Base + "/auth")
	{
		auth.POST(s.cfg.Gateway.Route.Endpoints["login"], s.handlers.HandlerLogin)
		auth.POST(s.cfg.Gateway.Route.Endpoints["register"], s.handlers.HandlerRegister)
	}

	r.Run(s.cfg.Gateway.Server.Port)
}

func (s *Server) authMiddleware(c *gin.Context) {

	jwt, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(401, gin.H{"error": "jwt not found"})
		c.Abort()
		return
	}

	resp, err := http.Get(handlers.StartURL + s.cfg.Authentication.Server.Port + s.cfg.Authentication.Route.Base + s.cfg.Authentication.Route.Endpoints["auth"] + "?token=" + jwt)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
