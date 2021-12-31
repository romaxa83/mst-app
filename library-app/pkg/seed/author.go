package seed

import (
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"gorm.io/gorm"
	"time"
)

func CreateAuthor(db *gorm.DB, name, country, bio string, birthday, death_date time.Time) error {
	return db.Create(&models.Author{
		Name:      name,
		Country:   country,
		Bio:       bio,
		Birthday:  birthday,
		DeathDate: death_date,
	}).Error
}
