package service

import (
	"authentication-service/internal/repository"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
)

type Service interface {
	Login(user common.AuthDataLogin) (common.Token, error)
	Register(user common.AuthDataRegister) error
	Logout(token common.Token) error
	ValidateToken(token common.Token) (error)
}

type ServiceManager struct {
	db repository.DateBase
}
