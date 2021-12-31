package seed

import (
	"gorm.io/gorm"
	"time"
)

func All() []Seed {

	l, err := time.LoadLocation("Europe/Vienna")
	if err != nil {
		panic(err)
	}

	return []Seed{
		Seed{
			Name: "Create Толстой",
			Run: func(db *gorm.DB) error {
				if err := CreateAuthor(
					db,
					"Толстой",
					"Российская империя",
					"биография",
					time.Date(2000, 1, 1, 0, 0, 0, 0, l),
					time.Date(2001, 1, 1, 0, 0, 0, 0, l),
				); err != nil {
					return err
				}

				return nil
			},
		},
		Seed{
			Name: "Create Пушкин",
			Run: func(db *gorm.DB) error {
				if err := CreateAuthor(
					db,
					"Пушкин",
					"Российская империя",
					"биография",
					time.Date(2000, 1, 1, 0, 0, 0, 0, l),
					time.Date(2001, 1, 1, 0, 0, 0, 0, l),
				); err != nil {
					return err
				}

				return nil
			},
		},
	}
}
