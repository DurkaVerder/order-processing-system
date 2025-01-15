package repository

import (
	"database/sql"
	"log"
	"os"

	common "github.com/DurkaVerder/common-for-order-processing-system/models"
	_ "github.com/lib/pq"
)

type DateBase interface {
	FindUser(user common.AuthDataLogin) (bool, error)
	AddUser(user common.AuthDataRegister) error
}

type RepositoryManager struct {
	db *sql.DB
}

func NewRepositoryManager() *RepositoryManager {
	return &RepositoryManager{
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
