package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/romaxa83/mst-app/gin-app/pkg/logger"
)

func NewClient(
	host,
	port,
	username,
	password,
	dbname,
	sslmode string,
) (*sqlx.DB, error) {
	// подключаемся к бд
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, username, dbname, password, sslmode))
	if err != nil {
		return nil, err
	}
	// пингуем подключение
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	logger.Info("Connect to db [postgres]")

	return db, nil
}
