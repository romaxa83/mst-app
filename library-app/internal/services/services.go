package services

import (
	"context"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
	"github.com/romaxa83/mst-app/library-app/pkg/storage"
	value_obj "github.com/romaxa83/mst-app/library-app/pkg/value-obj"
)

type Category interface {
	Create(input input.CreateCategory) (models.Category, error)
	GetAllPagination(query input.GetCategoryQuery) (db.Pagination, error)
	GetAllPaginationArchive(query input.GetCategoryQuery) (db.Pagination, error)
	GetAllList() ([]resources.CategoryListResource, error)
	GetOne(id int) (models.Category, error)
	Update(id int, input input.UpdateCategory) (models.Category, error)
	Delete(id int) error
	DeleteForce(id int) error
	Restore(id int) (models.Category, error)
}

type Author interface {
	Create(input input.CreateAuthor) (models.Author, error)
	GetAllPagination(query input.GetAuthorQuery) (db.Pagination, error)
	GetAllList() ([]resources.AuthorListResource, error)
	GetOne(id int) (models.Author, error)
	Update(id int, input input.UpdateAuthor) (models.Author, error)
	Delete(id int) error
}

type Book interface {
	Create(input input.CreateBook) (models.Book, error)
	GetAllPagination(query input.GetBookQuery) (db.Pagination, error)
	GetOne(id value_obj.ID) (models.Book, error)
	Update(id value_obj.ID, input input.UpdateBook) (models.Book, error)
	Delete(id value_obj.ID) error
}

type Media interface {
	UploadAndSaveFile(ctx context.Context, input input.UploadMedia) (models.Media, error)
}

type Services struct {
	Category Category
	Author   Author
	Book     Book
	Media    Media
}

type Deps struct {
	Repos           *repositories.Repo
	Environment     string
	StorageProvider storage.Provider
}

func NewServices(deps Deps) *Services {
	return &Services{
		Category: NewCategoryService(deps.Repos.Category),
		Author:   NewAuthorService(deps.Repos.Author),
		Book:     NewBookService(deps.Repos.Book),
		Media:    NewMediaService(deps.Repos.Media, deps.StorageProvider, deps.Environment),
	}
}
