package repository

type DataBase interface {
	GetUserEmailByOrderId(orderId int) (string, error)
}
