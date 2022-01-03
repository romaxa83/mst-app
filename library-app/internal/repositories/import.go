package repositories

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"gorm.io/gorm"
)

type ImportRepo struct {
	db *gorm.DB
}

func NewImportRepo(db *gorm.DB) *ImportRepo {
	return &ImportRepo{db}
}

func (r ImportRepo) Create(input input.CreateImport) (models.Import, error) {
	var model models.Import

	model.FilePath = input.FilePath
	model.Entity = input.Entity
	model.ContentType = input.ContentType
	model.Status = input.Status

	result := r.db.Create(&model)
	if result.Error != nil {
		return models.Import{}, result.Error
	}

	return model, nil
}

func (r ImportRepo) GetOneById(id int) (models.Import, error) {
	var model models.Import

	result := r.db.Find(&model, id).First(&model)
	if result.Error != nil {
		return model, result.Error
	}

	return model, nil
}
