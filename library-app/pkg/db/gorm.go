package db

import (
	"fmt"
	"github.com/romaxa83/mst-app/library-app/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient(host, port, user, password, dbname, sslmode string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}

	logger.Info("Connect to db [postgres-gorm]")

	return db, nil
}
