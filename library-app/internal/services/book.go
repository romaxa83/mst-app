package services

import (
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
	"github.com/romaxa83/mst-app/library-app/pkg/db"
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

func (s *BookService) GetOne(id int) (models.Book, error) {
	return s.repo.GetOneById(id)
}

func (s *BookService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *BookService) Update(id int, input input.UpdateBook) (models.Book, error) {
	return s.repo.Update(id, input)
}