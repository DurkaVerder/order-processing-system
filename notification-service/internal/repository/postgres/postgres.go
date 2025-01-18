package postgres

import (
	"database/sql"
	"log"
	"os"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	return &Postgres{db: initDb()}
}

func initDb() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	return db
}

func (p *Postgres) GetUserEmailByOrderId(orderId int) (string, error) {
	var email string
	err := p.db.QueryRow("SELECT email FROM users WHERE order_id=$1", orderId).Scan(&email)
	if err != nil {
		log.Printf("Failed to get user email: %s", err)
		return "", err
	}
	return email, nil
}
