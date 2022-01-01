package services

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
	value_obj "github.com/romaxa83/mst-app/library-app/pkg/value-obj"
)

type BookService struct {
	repo repositories.Book
}

func NewBookService(repo repositories.Book) *BookService {
	return &BookService{repo}
}

func (s *BookService) Create(input input.CreateBook) (models.Book, error) {
	return s.repo.Create(input)
}

func (s *BookService) GetAllPagination(query input.GetBookQuery) (db.Pagination, error) {
	return s.repo.GetAllPagination(query)
}

func (s *BookService) GetOne(id value_obj.ID) (models.Book, error) {
	return s.repo.GetOneById(id)
}

func (s *BookService) Delete(id value_obj.ID) error {
	return s.repo.Delete(id)
}

func (s *BookService) Update(id value_obj.ID, input input.UpdateBook) (models.Book, error) {
	return s.repo.Update(id, input)
}
