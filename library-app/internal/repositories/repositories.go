package repositories

//go:generate mockgen -source=repositories.go -destination=mocks/mock.go

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
	value_obj "github.com/romaxa83/mst-app/library-app/pkg/value-obj"
	"gorm.io/gorm"
)

type Category interface {
	Create(input input.CreateCategory) (models.Category, error)
	GetAllPagination(query input.GetCategoryQuery) (db.Pagination, error)
	GetAllPaginationArchive(query input.GetCategoryQuery) (db.Pagination, error)
	GetAllList() ([]resources.CategoryListResource, error)
	GetOneById(id int) (models.Category, error)
	Update(id int, input input.UpdateCategory) (models.Category, error)
	Delete(id int) error
	DeleteForce(id int) error
	Restore(id int) (models.Category, error)
}

type Author interface {
	Create(input input.CreateAuthor) (models.Author, error)
	GetAllPagination(query input.GetAuthorQuery) (db.Pagination, error)
	GetAllList() ([]resources.AuthorListResource, error)
	GetOneById(id int) (models.Author, error)
	GetOneByName(name string) (models.Author, error)
	Update(id int, input input.UpdateAuthor) (models.Author, error)
	Delete(id int) error
}

type Book interface {
	Create(input input.CreateBook) (models.Book, error)
	GetAllPagination(query input.GetBookQuery) (db.Pagination, error)
	GetAllList() ([]resources.BookListResource, error)
	GetOneById(id value_obj.ID) (models.Book, error)
	Update(id value_obj.ID, input input.UpdateBook) (models.Book, error)
	Delete(id value_obj.ID) error
}

type Media interface {
	Create(input input.UploadMedia) (models.Media, error)
	SetUrl(model models.Media, url string) (models.Media, error)
}

type Import interface {
	Create(input input.CreateImport) (models.Import, error)
	GetOneById(id int) (models.Import, error)
}

type Repo struct {
	Category
	Author
	Book
	Media
	Import
}

func NewRepositories(db *gorm.DB) *Repo {
	return &Repo{
		Category: NewCategoryRepo(db),
		Author:   NewAuthorRepo(db),
		Book:     NewBookRepo(db),
		Media:    NewMediaRepo(db),
		Import:   NewImportRepo(db),
	}
}
