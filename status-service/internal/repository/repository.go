package repository

type DataBase interface {
	UpdateStatus(orderID int, status string) error
	CreateRecordStatus(orderID int, status string) error
}
