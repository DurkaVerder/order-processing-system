package repository

import (
	common "github.com/DurkaVerder/common-for-order-processing-system/models"
)

// DateBase is an interface that contains methods for working with a database
type DateBase interface {
	FindUser(user common.AuthDataLogin) (int, string, error)
	AddUser(user common.AuthDataRegister) error
}

// Cache is an interface that contains methods for working with cache
type Cache interface {
	RevokeToken(token string) error
	IsTokenRevoked(token string) (bool, error)
}
