package models

import "gorm.io/gorm"

func InitModels(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&Author{},
		&Category{},
		&Book{},
		&Media{},
		&Import{},
	); err != nil {
		return err
	}

	return nil
}
