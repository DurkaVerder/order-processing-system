package repository

import (
	common "github.com/DurkaVerder/common-for-order-processing-system/models"
)

type DateBase interface {
	FindUser(user common.AuthDataLogin) (int, error)
	AddUser(user common.AuthDataRegister) error
}

type Cache interface {
	RevokeToken(token string) error
	IsTokenRevoked(token string) (bool, error)
}
