package repositories

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
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

func (r *CategoryRepo) GetAllPagination(pagination db.Pagination) (db.Pagination, error) {
	//logger.Infof("%+v", pagination)

	var categories []*models.Category

	r.db.Scopes(db.Paginate(categories, &pagination, r.db)).Find(&categories)
	pagination.Rows = categories

	return pagination, nil
}

type APIUser struct {
	Sort  int    `json:"title"`
	Title string `json:"sort"`
}

func (r *CategoryRepo) GetAllList() ([]resources.CategoryResource, error) {

	var resources []resources.CategoryResource
	result := r.db.Model(&models.Category{}).Find(&resources)
	if result.Error != nil {
		return resources, result.Error
	}

	return resources, nil
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
