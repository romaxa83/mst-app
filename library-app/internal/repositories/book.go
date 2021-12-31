package repositories

import (
	"fmt"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{db}
}

func (r *BookRepo) Create(input input.CreateBook) (models.Book, error) {

	var model models.Book
	model.Title = input.Title
	model.Desc = input.Desc
	model.AuthorID = input.AuthorID
	model.Pages = input.Pages
	model.Qty = input.Qty
	model.PublishedAt = input.PublishedAt

	result := r.db.Create(&model)
	if result.Error != nil {
		return models.Book{}, result.Error
	}

	return model, nil
}

func (r *BookRepo) GetOneById(id int) (models.Book, error) {

	var model models.Book

	result := r.db.Joins("Author").Find(&model, id).First(&model)
	if result.Error != nil {
		return model, result.Error
	}

	return model, nil
}

func (r *BookRepo) GetAllPagination(query input.GetBookQuery) (db.Pagination, error) {
	model := models.Book{}
	var models []models.Book
	var res []resources.BookResource
	pagination := query.Pagination

	q := r.db.Model(&model).Joins("Author")

	id := query.BookFilterQuery.Id
	if id != nil {
		q.Where("id = ?", *id)
	}

	search := query.Search.Search
	if search != "" {
		q.Where("name LIKE ?", fmt.Sprintf("%s%%", search))
	}

	q = q.Scopes(db.Paginate(&model, &pagination, r.db)).Find(&models)

	for _, m := range models {
		res = append(res, *resources.NewBookResource(m))
	}
	pagination.Rows = res

	return pagination, nil
}

func (r *BookRepo) Update(id int, input input.UpdateBook) (models.Book, error) {

	var model models.Book
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
	if nil != input.PublishedAt {
		model.PublishedAt = *input.PublishedAt
	}
	if nil != input.Pages {
		model.Pages = *input.Pages
	}
	if nil != input.Qty {
		model.Qty = *input.Qty
	}
	if nil != input.AuthorID {
		var a models.Author
		r.db.Find(&a, *input.AuthorID).First(&a)
		model.Author = a
	}
	r.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&input)

	return model, nil
}

func (r *BookRepo) Delete(id int) error {
	result := r.db.Delete(&models.Book{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
