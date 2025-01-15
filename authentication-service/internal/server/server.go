package server

import (
	"authentication-service/config"
	"authentication-service/internal/handlers"

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

	auth := r.Group(s.config.Authentication.Route.Base)
	{
		auth.POST(s.config.Authentication.Route.Endpoints["login"], s.handlers.Login)
		auth.POST(s.config.Authentication.Route.Endpoints["register"], s.handlers.Register)
		auth.GET(s.config.Authentication.Route.Endpoints["logout"], s.handlers.Logout)
		auth.GET(s.config.Authentication.Route.Endpoints["validate"], s.handlers.ValidateToken)
	}

	r.Run(s.config.Authentication.Server.Port)
}
