package services

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
)

type Category interface {
	Create(input input.CreateCategory) (models.Category, error)
	GetAllPagination(query input.GetCategoryQuery) (db.Pagination, error)
	GetAllList() ([]resources.CategoryResource, error)
	GetOne(id int) (models.Category, error)
	Update(id int, input input.UpdateCategory) (models.Category, error)
	Delete(id int) error
}

type Services struct {
	Category Category
}

type Deps struct {
	Repos *repositories.Repo
}

func NewServices(deps Deps) *Services {
	return &Services{
		Category: NewCategoryService(deps.Repos.Category),
	}
}
