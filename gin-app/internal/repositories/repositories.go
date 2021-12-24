package repositories

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"github.com/romaxa83/mst-app/gin-app/internal/repositories/postgres"
)

type Users interface {
	Create(ctx context.Context, user domains.User) (int, error)
	GetByCredentials(ctx context.Context, email, password string) (domains.User, error)
	GetByRefreshToken(ctx context.Context, refreshToken string) (domains.User, error)
	SetSession(ctx context.Context, userId int, session domains.Session) error
}

type TodoList interface {
	Create(ctx context.Context, userId int, list domains.TodoList) (int, error)
	GetAll(userId int) ([]domains.TodoList, error)
	GetById(userId, lisId int) (domains.TodoList, error)
	Update(userId, lisId int, input domains.UpdateTodoListInput) error
	Delete(userId, lisId int) error
}

type TodoItem interface {
	Create(listId int, item domains.TodoItem) (int, error)
	GetAll(userId, listId int) ([]domains.TodoItem, error)
	GetById(userId, itemId int) (domains.TodoItem, error)
	Update(userId, itemId int, input domains.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type Files interface {
	Create(ctx context.Context, file domains.File) (int, error)
	//UpdateStatus(ctx context.Context, fileName string, status domain.FileStatus) error
	//GetForUploading(ctx context.Context) (domain.File, error)
	//UpdateStatusAndSetURL(ctx context.Context, id primitive.ObjectID, url string) error
	//GetByID(ctx context.Context, id, schoolID primitive.ObjectID) (domain.File, error)
}

type Repositories struct {
	Users
	TodoList
	TodoItem TodoItem
	Files
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Users:    postgres.NewUserRepo(db),
		TodoList: postgres.NewTodoListRepo(db),
		TodoItem: postgres.NewTodoItemRepo(db),
		Files:    postgres.NewFilesRepo(db),
	}
}
