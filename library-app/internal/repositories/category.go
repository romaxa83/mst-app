package repositories

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db}
}

func (r *CategoryRepo) Create(input input.CreateCategory) (models.Category, error) {

	var model models.Category
	model.Title = input.Title
	model.Desc = input.Desc

	result := r.db.Create(&model)
	if result.Error != nil {
		return models.Category{}, result.Error
	}

	return model, nil
}

func (r *CategoryRepo) GetAll() ([]models.Category, error) {

	var models []models.Category

	result := r.db.Find(&models)
	if result.Error != nil {
		return models, result.Error
	}

	return models, nil
}

func (r *CategoryRepo) GetOneById(id int) (models.Category, error) {

	var model models.Category

	result := r.db.Find(&model, id).First(&model)
	if result.Error != nil {
		return model, result.Error
	}

	return model, nil
}

func (r *CategoryRepo) Update(id int, input input.UpdateCategory) (models.Category, error) {

	var model models.Category
	model, err := r.GetOneById(id)
	if err != nil {
		return model, err
	}

	if nil != input.Title {
		model.Title = *input.Title
	}
	if nil != input.Desc {
		model.Desc = *input.Desc
	}
	if nil != input.Sort {
		model.Sort = *input.Sort
	}
	if nil != input.Active {
		model.Active = *input.Active
	}
	r.db.Save(&model)

	return model, nil
}

func (r *CategoryRepo) Delete(id int) error {

	result := r.db.Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
