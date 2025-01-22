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

func (p *Postgres) UpdateStatus(orderID int, status string) error {
	_, err := p.db.Exec(updateStatusQuery, status, orderID)
	return err
}

func (p *Postgres) CreateRecordStatus(orderID int, status string) error {
	_, err := p.db.Exec(createRecordStatusQuery, orderID, status)
	return err
}
