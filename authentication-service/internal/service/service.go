package service

import (
	"authentication-service/internal/jwt"
	"authentication-service/internal/kafka"
	"authentication-service/internal/kafka/producer"
	"authentication-service/internal/repository"
	"errors"
	"log"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(user common.AuthDataLogin) (common.Token, error)
	Register(user common.AuthDataRegister) error
	Logout(token common.Token) error
	ValidateToken(token common.Token) error
}

type ServiceManager struct {
	db       repository.DateBase
	cache    repository.Cache
	producer producer.Producer
}

func NewServiceManager(db repository.DateBase, cache repository.Cache, producer producer.Producer) *ServiceManager {
	return &ServiceManager{
		db:       db,
		cache:    cache,
		producer: producer,
	}
}

func (s *ServiceManager) Login(user common.AuthDataLogin) (common.Token, error) {
	hashedPassword, err := s.hashPassword(user.Password)
	if err != nil {
		return common.Token{}, err
	}
	user.Password = hashedPassword

	userId, err := s.db.FindUser(user)
	if err != nil {
		return common.Token{}, err
	}

	token, err := jwt.GenerateToken(userId)
	if err != nil {
		return common.Token{}, err
	}

	return token, nil
}

func (s *ServiceManager) Register(user common.AuthDataRegister) error {
	hashedPassword, err := s.hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	err = s.db.AddUser(user)
	if err != nil {
		return err
	}

	go func() {
		if err := s.sendMessages(user); err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}
		log.Printf("Message sent to %s", user.Email)
	}()

	return nil
}

func (s *ServiceManager) Logout(token common.Token) error {
	err := s.cache.RevokeToken(token.Token)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceManager) ValidateToken(token common.Token) error {
	isRevoked, err := s.cache.IsTokenRevoked(token.Token)
	if err != nil {
		return err
	}

	if isRevoked {
		return errors.New("token is revoked")
	}

	return nil
}

func (s *ServiceManager) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil

}

func (s *ServiceManager) sendMessages(data common.AuthDataRegister) error {
	notification := s.createRegisterNotification(data)

	if err := s.producer.SendMessage(kafka.NotificationTopic, notification); err != nil {
		return err
	}

	return nil
}

func (s *ServiceManager) createRegisterNotification(data common.AuthDataRegister) common.Notification {
	return common.Notification{
		To:      data.Email,
		Subject: "Registration",
		Body:    "You have successfully registered",
	}
}
