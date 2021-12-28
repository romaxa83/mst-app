package repositories

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
	"gorm.io/gorm"
)

type Category interface {
	Create(input input.CreateCategory) (models.Category, error)
	GetAllPagination(pagination db.Pagination) (db.Pagination, error)
	GetAllList() ([]resources.CategoryResource, error)
	GetOneById(id int) (models.Category, error)
	Update(id int, input input.UpdateCategory) (models.Category, error)
	Delete(id int) error
}

type Repo struct {
	Category
}

func NewRepositories(db *gorm.DB) *Repo {
	return &Repo{
		Category: NewCategoryRepo(db),
	}
}
