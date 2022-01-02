package repositories

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"gorm.io/gorm"
)

type MediaRepo struct {
	db *gorm.DB
}

func NewMediaRepo(db *gorm.DB) *MediaRepo {
	return &MediaRepo{db}
}

func (r MediaRepo) Create(input input.UploadMedia) (models.Media, error) {
	var model models.Media
	model.OwnerID = input.OwnerID
	model.OwnerType = input.OwnerType
	model.Type = input.Type
	model.ContentType = input.ContentType
	model.Name = input.Name
	model.Size = input.Size

	result := r.db.Create(&model)
	if result.Error != nil {
		return models.Media{}, result.Error
	}

	return model, nil
}

func (r MediaRepo) SetUrl(model models.Media, url string) (models.Media, error) {
	result := r.db.Model(&model).Updates(models.Media{
		URL:    url,
		Status: models.UploadedToStorage,
	})
	if result.Error != nil {
		return model, result.Error
	}

	return model, nil
}
