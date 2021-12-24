package services

import (
	"context"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"github.com/romaxa83/mst-app/gin-app/internal/repositories"
	"time"
)

type TodoListService struct {
	repo repositories.TodoList
}

func NewTodoListService(repo repositories.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(ctx context.Context, userId int, list domains.TodoList) (int, error) {
	l := domains.TodoList{
		Title:       list.Title,
		Description: list.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return s.repo.Create(ctx, userId, l)
}

func (s *TodoListService) GetAll(ctx context.Context, userId int) ([]domains.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(ctx context.Context, userId, listId int) (domains.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(ctx context.Context, userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListService) Update(ctx context.Context, userId, listId int, input domains.UpdateTodoListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}
