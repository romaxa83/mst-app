package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Title  string `json:"title" binding:"required,max=256" gorm:"size:256"`
	Desc   string `json:"desc"`
	Active bool   `gorm:"default:true"`
	Sort   int    `gorm:"default:0"`
}
