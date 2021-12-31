package db

import (
	"gorm.io/gorm"
)

func Paginate(value interface{}, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.TotalRows = totalRows

	// todo.txt доделать коректное отображение totalPages
	var totalPages int
	if float64(totalRows) <= float64(pagination.Limit) {
		totalPages = 1
	} else {
		totalPages = int(float64(totalRows) / float64(pagination.Limit))
	}

	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}
