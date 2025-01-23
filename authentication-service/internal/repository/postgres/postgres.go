package postgres

import (
	"database/sql"
	"log"
	"os"
	"time"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	_ "github.com/lib/pq"
)

// Postgres is a struct for working with PostgreSQL
type Postgres struct {
	db *sql.DB
}

// NewPostgres creates a new Postgres struct
func NewPostgres() *Postgres {
	return &Postgres{
		db: initDb(),
	}
}

// initDb initializes a connection to the database
func initDb() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	return db
}

// FindUser finds a user in the database
func (p *Postgres) FindUser(user common.AuthDataLogin) (int, string, error) {
	var userId int
	var password string
	err := p.db.QueryRow(findUserQuery, user.Login).Scan(&userId, &password)
	if err != nil {
		return -1, "", err
	}

	return userId, password, nil
}

// AddUser adds a user to the database
func (p *Postgres) AddUser(user common.AuthDataRegister) error {
	time := time.Now()
	_, err := p.db.Exec(addUserQuery, user.Email, user.Login, user.Password, user.Username, time, time)
	if err != nil {
		return err
	}

	return nil
}
