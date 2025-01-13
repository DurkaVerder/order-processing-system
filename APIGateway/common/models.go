package common

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type User struct {
	Id        int       `json:"id,omitempty"`
	Email     string    `json:"email"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"jwt"`
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

type Order struct {
	Id          int       `json:"id,omitempty"`
	UserId      int       `json:"user_id"`
	TotalAmount int       `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at"`
}

type OrderItem struct {
	Id          int       `json:"id,omitempty"`
	OrderId     int       `json:"order_id"`
	ProductName string    `json:"product_name"`
	Quantity    int       `json:"quantity"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

type StatusOrder struct {
	Id        int       `json:"id,omitempty"`
	OrderId   int       `json:"order_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type HistoryOrder struct {
	Id        int           `json:"id,omitempty"`
	OrderId   int           `json:"order_id"`
	List      []StatusOrder `json:"list"`
	CreatedAt time.Time     `json:"created_at"`
}
