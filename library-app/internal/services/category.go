package services

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
)

type CategoryService struct {
	repo repositories.Category
}

func NewCategoryService(repo repositories.Category) *CategoryService {
	return &CategoryService{repo}
}

func (s *CategoryService) Create(input input.CreateCategory) (models.Category, error) {
	return s.repo.Create(input)
}

func (s *CategoryService) GetAllPagination(query input.GetCategoryQuery) (db.Pagination, error) {
	return s.repo.GetAllPagination(query)
}

func (s *CategoryService) GetAllList() ([]resources.CategoryResource, error) {
	return s.repo.GetAllList()
}

func (s *CategoryService) GetOne(id int) (models.Category, error) {
	return s.repo.GetOneById(id)
}

func (s *CategoryService) Update(id int, input input.UpdateCategory) (models.Category, error) {
	return s.repo.Update(id, input)
}

func (s *CategoryService) Delete(id int) error {
	return s.repo.Delete(id)
}
