package repositories

import (
	"fmt"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
	"github.com/romaxa83/mst-app/library-app/pkg/logger"
	"gorm.io/gorm"
)

type AuthorRepo struct {
	db *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) *AuthorRepo {
	return &AuthorRepo{db}
}

func (r *AuthorRepo) Create(input input.CreateAuthor) (models.Author, error) {

	var model models.Author
	model.Name = input.Name
	model.Country = input.Country
	model.Bio = input.Bio
	model.Birthday = input.Birthday
	model.DeathDate = input.DeathDate

	result := r.db.Create(&model)
	if result.Error != nil {
		return models.Author{}, result.Error
	}

	return model, nil
}

func (r *AuthorRepo) GetAllPagination(query input.GetAuthorQuery) (db.Pagination, error) {
	model := models.Author{}
	pagination := query.Pagination
	var resources []resources.AuthorResource
	q := r.db.Model(&model)

	id := query.AuthorFilterQuery.Id
	if id != nil {
		q.Where("id = ?", *id)
	}

	search := query.Search.Search
	if search != "" {
		q.Where("name LIKE ?", fmt.Sprintf("%s%%", search))
	}

	q = q.Scopes(db.Paginate(&model, &pagination, r.db)).Find(&resources)
	logger.Warnf("%+v", resources)
	pagination.Rows = resources

	return pagination, nil
}

func (r *AuthorRepo) GetAllList() ([]resources.AuthorResource, error) {

	var resources []resources.AuthorResource
	result := r.db.Model(&models.Author{}).Find(&resources)
	if result.Error != nil {
		return resources, result.Error
	}

	return resources, nil
}

func (r *AuthorRepo) GetOneById(id int) (models.Author, error) {

	var model models.Author

	result := r.db.Find(&model, id).First(&model)
	if result.Error != nil {
		return model, result.Error
	}

	return model, nil
}

func (r *AuthorRepo) Update(id int, input input.UpdateAuthor) (models.Author, error) {

	var model models.Author
	model, err := r.GetOneById(id)
	if err != nil {
		return model, err
	}

	if nil != input.Name {
		model.Name = *input.Name
	}
	if nil != input.Country {
		model.Country = *input.Country
	}
	if nil != input.Bio {
		model.Bio = *input.Bio
	}
	if nil != input.Birthday {
		model.Birthday = *input.Birthday
	}
	if nil != input.DeathDate {
		model.DeathDate = *input.DeathDate
	}
	r.db.Save(&model)

	return model, nil
}

func (r *AuthorRepo) Delete(id int) error {

	result := r.db.Delete(&models.Author{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
