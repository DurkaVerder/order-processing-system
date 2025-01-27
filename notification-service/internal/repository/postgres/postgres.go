package postgres

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	return &Postgres{db: initDb()}
}

func initDb() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	return db
}

func (p *Postgres) GetUserEmailByOrderId(orderId int) (string, error) {
	var email string
	err := p.db.QueryRow(SelectUserEmailByOrderId, orderId).Scan(&email)
	if err != nil {
		log.Printf("Failed to get user email: %s", err)
		return "", err
	}
	return email, nil
}
