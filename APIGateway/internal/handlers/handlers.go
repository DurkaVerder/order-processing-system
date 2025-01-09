package handlers

import "APIGateway/config"

// Handlers is an interface that defines the methods that the handlers must implement
type Handlers interface {
}

type HandlersManager struct {
	cfg config.Config
}

func NewHandlersManager(cfg config.Config) Handlers {
	return &HandlersManager{
		cfg: cfg,
	}
}
