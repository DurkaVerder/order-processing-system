package server

import (
	"APIGateway/config"
	"APIGateway/internal/handlers"
	"errors"
	"net/http"
	"os"
	"strings"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Server is a struct that represents the server
type Server struct {
	handlers handlers.Handlers
	cfg      config.Config
}

// NewServer is a constructor for Server
func NewServer(handlers handlers.Handlers, cfg config.Config) *Server {
	return &Server{
		handlers: handlers,
		cfg:      cfg,
	}
}

// Start starts the server
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
		protected.GET(s.cfg.Gateway.Route.Endpoints["history_order"], s.handlers.HandlerHistoryOrder)
	}
	auth := r.Group(s.cfg.Gateway.Route.Base + "/auth")
	{
		auth.POST(s.cfg.Gateway.Route.Endpoints["login"], s.handlers.HandlerLogin)
		auth.POST(s.cfg.Gateway.Route.Endpoints["register"], s.handlers.HandlerRegister)
	}

	r.Run(s.cfg.Gateway.Server.Port)
}

// authMiddleware is a middleware for authentication
func (s *Server) authMiddleware(c *gin.Context) {

	jwt, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(401, gin.H{"error": "jwt not found"})
		c.Abort()
		return
	}

	resp, err := http.Get(handlers.StartURLauth + s.cfg.Authentication.Server.Port + s.cfg.Authentication.Route.Base + s.cfg.Authentication.Route.Endpoints["validate"] + "?token=" + jwt)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	userId, err := GetUserIdFromToken(jwt, os.Getenv("SECRET_KEY"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	// Refresh token
	c.SetCookie("jwt", jwt, 3600*72, "/", "", false, true)
	c.Set("user_id", userId)
	c.Next()
}

// GetUserIdFromToken returns user id from token
func GetUserIdFromToken(tokenString string, secretKey string) (int, error) {
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return 0, errors.New("invalid token format")
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.ParseWithClaims(tokenString, &common.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*common.Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, errors.New("invalid token")
}
