package postgres

import (
	"database/sql"
	"log"
	"os"
	"time"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	return &Postgres{
		db: initDb(),
	}
}

func initDb() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	return db
}

func (p *Postgres) FindUser(user common.AuthDataLogin) (int, error) {
	var userId int
	err := p.db.QueryRow(findUserQuery, user.Login, user.Password).Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, nil
}

func (p *Postgres) AddUser(user common.AuthDataRegister) error {
	time := time.Now()
	_, err := p.db.Exec(addUserQuery, user.Email, user.Login, user.Password, user.Username, time, time)
	if err != nil {
		return err
	}

	return nil
}
