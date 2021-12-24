package services

import (
	"context"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"github.com/romaxa83/mst-app/gin-app/internal/repositories"
	"time"
)

type TodoItemService struct {
	repo     repositories.TodoItem
	listRepo repositories.TodoList
}

func NewTodoItemService(repo repositories.TodoItem, listRepo repositories.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(ctx context.Context, userId, listId int, item domains.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}

	i := domains.TodoItem{
		Title:       item.Title,
		Description: item.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return s.repo.Create(listId, i)
}

func (s *TodoItemService) GetAll(ctx context.Context, userId, listId int) ([]domains.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(ctx context.Context, userId, itemId int) (domains.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input domains.UpdateItemInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, itemId, input)
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}
