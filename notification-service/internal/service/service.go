package service

import (
	"context"
	"fmt"
	"log"
	"notification-service/internal/repository"
	"os"
	"sync"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	"gopkg.in/gomail.v2"
)

type Service interface {
	StarProcessingAndSendingMsg(numWorkers int, ctx context.Context)
	StopProcessingAndSendingMsg()
	AddDataForNotifyInChan(n common.DataForNotify)
}

type ServiceManager struct {
	db            repository.DataBase
	dataForNotify chan common.DataForNotify
	mailQueue     chan common.Notification
	wg            sync.WaitGroup
}

func NewServiceManager(db repository.DataBase) *ServiceManager {
	return &ServiceManager{
		db:            db,
		dataForNotify: make(chan common.DataForNotify, 100),
		mailQueue:     make(chan common.Notification, 100),
		wg:            sync.WaitGroup{},
	}
}

func (s *ServiceManager) CreateNotification(notify common.DataForNotify) error {
	notification := common.Notification{}
	switch notify.Event {
	case "order_created":
		notification.Subject = "Order created"
		notification.Body = "Order has been created"

	case "order_updated":
		notification.Subject = "Order updated"
		notification.Body = "Order has been updated"
	case "register":
		notification.Subject = "Register"
		notification.Body = "Welcome!"
		notification.To = notify.UserEmail
	}
	if notification.To == "" {
		email, err := s.db.GetUserEmailByOrderId(notify.OrderId)
		if err != nil {
			log.Printf("Error get email: %s", err)
			return err
		}
		notification.To = email
	}

	select {
	case s.mailQueue <- notification:
		log.Println("msg send in queue")
	default:
		log.Printf("Mail queue is full, notification dropped")
		return fmt.Errorf("mail queue is full")
	}
	return nil
}

func (s *ServiceManager) SendNotification(notify common.Notification) error {

	msg := gomail.NewMessage()
	from := os.Getenv("MAIL")
	msg.SetHeader("From", from)
	msg.SetHeader("To", notify.To)
	msg.SetHeader("Subject", notify.Subject)
	msg.SetBody("text/plain", notify.Body)

	d := gomail.NewDialer("smtp.mail.ru", 465, from, os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(msg); err != nil {
		log.Printf("Error send msg %s", err)
		return err
	}
	log.Printf("Send msg to: %s", notify.To)

	return nil
}

func (s *ServiceManager) workerToSendMsg(ctx context.Context) {
	defer s.wg.Done()
	for {
		select {
		case notify, ok := <-s.mailQueue:
			if !ok {
				log.Println("Worker stopping: mail queue closed")
				return
			}
			if err := s.SendNotification(notify); err != nil {
				log.Printf("Error sending notification: %s", err)
			}
		case <-ctx.Done():
			log.Println("Worker stopping: context cancelled")
			return
		}

	}
}

func (s *ServiceManager) workerToCreateMsg(ctx context.Context) {
	defer s.wg.Done()
	for {
		select {
		case data, ok := <-s.dataForNotify:
			if !ok {
				log.Println("Worker stopping: data mail queue closed")
				return
			}
			if err := s.CreateNotification(data); err != nil {
				log.Printf("Error create notification: %s", err)
			}
		case <-ctx.Done():
			log.Println("Worker stopping: context cancelled")
			return
		}

	}
}

func (s *ServiceManager) StarProcessingAndSendingMsg(numWorkers int, ctx context.Context) {
	for i := 0; i < numWorkers; i++ {
		s.wg.Add(1)
		go s.workerToCreateMsg(ctx)
		s.wg.Add(1)
		go s.workerToSendMsg(ctx)
	}
}

func (s *ServiceManager) StopProcessingAndSendingMsg() {
	close(s.dataForNotify)
	close(s.mailQueue)
	s.wg.Wait()
	log.Println("All workers stopped")
}

func (s *ServiceManager) AddDataForNotifyInChan(n common.DataForNotify) {
	s.dataForNotify <- n
}
