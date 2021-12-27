package services

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
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

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
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
