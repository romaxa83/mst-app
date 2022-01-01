package repositories

import (
	"fmt"
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

func (r *CategoryRepo) GetAllPagination(query input.GetCategoryQuery) (db.Pagination, error) {
	category := models.Category{}
	pagination := query.Pagination
	var resources []*resources.CategoryResource

	q := r.db.Model(&category)

	id := query.CategoryFilterQuery.Id
	if id != nil {
		q.Where("id = ?", *id)
	}

	active := query.CategoryFilterQuery.Active
	if active != nil {
		q.Where("active = ?", *active)
	}

	sort := query.CategoryFilterQuery.Sort
	if sort != nil {
		q.Where("sort = ?", *sort)
	}

	search := query.Search.Search
	if search != "" {
		q.Where("title LIKE ?", fmt.Sprintf("%s%%", search))
	}

	q = q.Scopes(db.Paginate(&category, &pagination, r.db)).Find(&resources)

	pagination.Rows = resources

	return pagination, nil
}

func (r *CategoryRepo) GetAllPaginationArchive(query input.GetCategoryQuery) (db.Pagination, error) {
	category := models.Category{}
	pagination := query.Pagination
	var resources []*resources.CategoryResource

	q := r.db.Unscoped().Model(&category).Where("deleted_at IS NOT NULL")

	id := query.CategoryFilterQuery.Id
	if id != nil {
		q.Where("id = ?", *id)
	}

	active := query.CategoryFilterQuery.Active
	if active != nil {
		q.Where("active = ?", *active)
	}

	sort := query.CategoryFilterQuery.Sort
	if sort != nil {
		q.Where("sort = ?", *sort)
	}

	search := query.Search.Search
	if search != "" {
		q.Where("title LIKE ?", fmt.Sprintf("%s%%", search))
	}

	q = q.Scopes(db.Paginate(&category, &pagination, r.db)).Find(&resources)

	pagination.Rows = resources

	return pagination, nil
}

func (r *CategoryRepo) GetAllList() ([]resources.CategoryListResource, error) {

	var resources []resources.CategoryListResource
	result := r.db.Model(&models.Category{}).Scopes(db.Active).Find(&resources)
	if result.Error != nil {
		return resources, result.Error
	}

	return resources, nil
}

func (r *CategoryRepo) GetOneById(id int) (models.Category, error) {

	var model models.Category

	result := r.db.Find(&model, id).Preload("Books").First(&model)
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

func (r *CategoryRepo) Restore(id int) (models.Category, error) {

	var model models.Category

	result := r.db.Unscoped().Find(&model, id).First(&model)
	if result.Error != nil {
		return model, result.Error
	}

	r.db.Unscoped().Model(&model).UpdateColumn("deleted_at", nil)

	return model, nil
}

func (r *CategoryRepo) DeleteForce(id int) error {

	result := r.db.Unscoped().Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
