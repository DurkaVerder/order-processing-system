package postgres

import (
	"database/sql"
	"log"
	"os"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	return &Postgres{
		db: initDB(),
	}
}

func initDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error while connecting to DB: %v", err)
	}

	return db
}

func (p *Postgres) AddOrder(order common.Order) error {
	_, err := p.db.Exec(addOrderQuery, order.UserId, order.TotalAmount, order.Status, order.CreatedAt, order.UpdateAt)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetOrder(id string) (common.Order, error) {
	var order common.Order
	err := p.db.QueryRow(getOrderQuery, id).Scan(&order.Id, &order.UserId, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.UpdateAt)
	if err != nil {
		return order, err
	}

	return order, nil
}

func (p *Postgres) GetAllOrders(userId string) ([]common.Order, error) {
	var orders []common.Order
	rows, err := p.db.Query(getAllOrdersQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order common.Order
		err := rows.Scan(&order.Id, &order.UserId, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.UpdateAt)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func (p *Postgres) DeleteOrder(id string) error {
	_, err := p.db.Exec(deleteOrderQuery, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUserEmail(userId string) (string, error) {
	var email string
	err := p.db.QueryRow(GetUserEmailQuery, userId).Scan(&email)
	if err != nil {
		return "", err
	}

	return email, nil
}
