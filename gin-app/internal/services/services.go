package services

import (
	"context"
	"github.com/romaxa83/mst-app/gin-app/internal/config"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"github.com/romaxa83/mst-app/gin-app/internal/repositories"
	"github.com/romaxa83/mst-app/gin-app/pkg/auth"
	"github.com/romaxa83/mst-app/gin-app/pkg/email"
	"github.com/romaxa83/mst-app/gin-app/pkg/hash"
	"github.com/romaxa83/mst-app/gin-app/pkg/otp"
	"github.com/romaxa83/mst-app/gin-app/pkg/storage"
	"io"
	"time"
)

//go:generate mockgen -source=services.go -destination=mocks/mock.go

type UserSignUpInput struct {
	Name     string `json:"name" binding:"required,min=2,max=64"`
	Email    string `json:"email" binding:"required,email,max=64"`
	Phone    string `json:"phone" binding:"required,max=13"`
	Password string `json:"password" binding:"required,min=8,max=64"`
}

type UserSignInInput struct {
	Email    string
	Password string
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Users interface {
	SignUp(ctx context.Context, input UserSignUpInput) (int, error)
	SignIn(ctx context.Context, input UserSignInInput) (Tokens, error)
	RefreshTokens(ctx context.Context, refreshToken string) (Tokens, error)
}

type TodoList interface {
	Create(ctx context.Context, userId int, list domains.TodoList) (int, error)
	GetAll(ctx context.Context, userId int) ([]domains.TodoList, error)
	GetById(ctx context.Context, userId, listId int) (domains.TodoList, error)
	Update(ctx context.Context, userId, listId int, input domains.UpdateTodoListInput) error
	Delete(ctx context.Context, userId, listId int) error
}

type TodoItem interface {
	Create(ctx context.Context, userId, listId int, item domains.TodoItem) (int, error)
	GetAll(ctx context.Context, userId, listId int) ([]domains.TodoItem, error)
	GetById(ctx context.Context, userId, itemId int) (domains.TodoItem, error)
	Update(userId, itemId int, input domains.UpdateItemInput) error
	Delete(userId, itemId int) error
}

type VerificationEmailInput struct {
	Email            string
	Name             string
	VerificationCode string
}

type Emails interface {
	SendVerificationEmail(VerificationEmailInput) error
}

type UploadInput struct {
	File        io.Reader
	Filename    string
	Size        int64
	ContentType string
	Type        domains.FileType
}

type Files interface {
	UploadAndSaveFile(ctx context.Context, file domains.File) (string, error)
	Save(ctx context.Context, file domains.File) (int, error)
	//UpdateStatus(ctx context.Context, fileName string, status domain.FileStatus) error // TODO check schoolID
	//GetByID(ctx context.Context, id, schoolId primitive.ObjectID) (domain.File, error)
	//InitStorageUploaderWorkers(ctx context.Context)
}

type Services struct {
	Users    Users
	TodoList TodoList
	TodoItem TodoItem
	Files    Files
}

type Dependencies struct {
	Repos                  *repositories.Repositories
	Hasher                 hash.PasswordHasher
	TokenManager           auth.TokenManager
	AccessTokenTTL         time.Duration
	RefreshTokenTTL        time.Duration
	EmailConfig            config.EmailConfig
	EmailSender            email.Sender
	OtpGenerator           otp.Generator
	VerificationCodeLength int
	StorageProvider        storage.Provider
	Environment            string
}

func NewServices(dependencies Dependencies) *Services {

	emailsService := NewEmailsService(dependencies.EmailConfig, dependencies.EmailSender)

	fileService := NewFilesService(
		dependencies.Repos.Files,
		dependencies.StorageProvider,
		dependencies.Environment,
	)

	usersService := NewUsersService(
		dependencies.Repos.Users,
		dependencies.Hasher,
		dependencies.TokenManager,
		dependencies.AccessTokenTTL,
		dependencies.RefreshTokenTTL,
		emailsService,
		dependencies.OtpGenerator,
		dependencies.VerificationCodeLength,
	)

	return &Services{
		Users:    usersService,
		TodoList: NewTodoListService(dependencies.Repos.TodoList),
		TodoItem: NewTodoItemService(dependencies.Repos.TodoItem, dependencies.Repos.TodoList),
		Files:    fileService,
	}
}
