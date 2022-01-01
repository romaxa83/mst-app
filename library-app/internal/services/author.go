package services

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/resources"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
)

type AuthorService struct {
	repo repositories.Author
}

func NewAuthorService(repo repositories.Author) *AuthorService {
	return &AuthorService{repo}
}

func (s *AuthorService) Create(input input.CreateAuthor) (models.Author, error) {
	return s.repo.Create(input)
}

func (s *AuthorService) GetAllPagination(query input.GetAuthorQuery) (db.Pagination, error) {
	return s.repo.GetAllPagination(query)
}

func (s *AuthorService) GetAllList() ([]resources.AuthorListResource, error) {
	return s.repo.GetAllList()
}

func (s *AuthorService) GetOne(id int) (models.Author, error) {
	return s.repo.GetOneById(id)
}

func (s *AuthorService) Update(id int, input input.UpdateAuthor) (models.Author, error) {
	return s.repo.Update(id, input)
}

func (s *AuthorService) Delete(id int) error {
	return s.repo.Delete(id)
}
