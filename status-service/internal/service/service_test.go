package service

import (
	"errors"
	"testing"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDBRepository struct {
	mock.Mock
}

func (m *MockDBRepository) UpdateStatus(orderId int, status string) error {
	args := m.Called(orderId, status)
	return args.Error(0)
}

func (m *MockDBRepository) CreateRecordStatus(orderId int, status string) error {
	args := m.Called(orderId, status)
	return args.Error(0)
}

type MockProducer struct {
	mock.Mock
}

func (m *MockProducer) SendMessage(topic string, message common.DataForNotify) error {
	args := m.Called(topic, message)
	return args.Error(0)
}

func TestChangeStatus(t *testing.T) {
	mockDB := MockDBRepository{}

	mockProducer := MockProducer{}

	mockDB.On("UpdateStatus", 1, "in-processing").Return(nil)
	mockDB.On("CreateRecordStatus", 1, "in-processing").Return(nil)
	mockProducer.On("SendMessage", "notification", common.DataForNotify{
		Event:   "update_status",
		OrderId: 1,
		Status:  "in-processing",
	}).Return(nil)

	service := NewServiceManager(&mockDB, &mockProducer)

	err := service.ChangeStatus(1, "in-processing")
	assert.NoError(t, err)

	mockDB.AssertExpectations(t)
}

func TestChangeStatus_UpdateStatusError(t *testing.T) {
	mockDB := MockDBRepository{}
	mockProducer := MockProducer{}

	mockDB.On("UpdateStatus", 1, "in-processing").Return(errors.New("error connection to db"))

	service := NewServiceManager(&mockDB, &mockProducer)

	err := service.ChangeStatus(1, "in-processing")
	assert.EqualError(t, err, "error connection to db")

	mockDB.AssertExpectations(t)
	mockProducer.AssertNotCalled(t, "SendMessage")
}

func TestChangeStatus_CreateRecordStatusError(t *testing.T) {
	mockDB := MockDBRepository{}
	mockProducer := MockProducer{}

	mockDB.On("UpdateStatus", 1, "in-processing").Return(nil)
	mockDB.On("CreateRecordStatus", 1, "in-processing").Return(errors.New("error connection to db"))

	service := NewServiceManager(&mockDB, &mockProducer)

	err := service.ChangeStatus(1, "in-processing")
	assert.EqualError(t, err, "error connection to db")

	mockDB.AssertExpectations(t)
	mockProducer.AssertNotCalled(t, "SendMessage")
}
func TestChangeStatusInvalidData(t *testing.T) {
	mockDB := MockDBRepository{}
	mockProducer := MockProducer{}

	service := NewServiceManager(&mockDB, &mockProducer)

	err := service.ChangeStatus(-1, "in-processing")
	assert.EqualError(t, err, "invalid orderId")

	err = service.ChangeStatus(1, "")
	assert.EqualError(t, err, "invalid status")

	mockDB.AssertNotCalled(t, "UpdateStatus")
	mockDB.AssertNotCalled(t, "CreateRecordStatus")
	mockProducer.AssertNotCalled(t, "SendMessage")
}

func TestCreateMessage(t *testing.T) {
	service := ServiceManager{}

	notify := service.createMessage(1, "in-processing")
	assert.Equal(t, common.DataForNotify{
		Event:   "update_status",
		OrderId: 1,
		Status:  "in-processing",
	}, notify)

	notify = service.createMessage(2, "completed")
	assert.Equal(t, common.DataForNotify{
		Event:   "update_status",
		OrderId: 2,
		Status:  "completed",
	}, notify)
}
