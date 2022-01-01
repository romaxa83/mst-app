package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	AuthorID    int
	Author      Author     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Categories  []Category `gorm:"many2many:book_category;"`
	Title       string     `json:"title" binding:"required,max=256" gorm:"size:256"`
	Desc        string     `json:"desc"`
	Active      bool       `gorm:"default:true"`
	Sort        int        `gorm:"default:0"`
	PublishedAt time.Time  `gorm:"not null"`
	Pages       int        `gorm:"not null"`
	Qty         int        `gorm:"default:0"`
}
