package models

import "gorm.io/gorm"

func InitModels(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&Author{},
		&Category{},
		&Book{},
	); err != nil {
		return err
	}

	return nil
}
