package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	userTable       = "users"
	todoListTable   = "todo_list"
	userListTable   = "user_list"
	todoItemTable   = "todo_items"
	listsItemsTable = "list_items"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Pinged success")

	return db, nil
}
