package models

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	gorm.Model
	Name      string `gorm:"size:256"`
	Country   string `gorm:"size:256"`
	Birthday  time.Time
	DeathDate time.Time
	Bio       string
}
