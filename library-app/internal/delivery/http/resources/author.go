package resources

import (
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"time"
)

type AuthorResource struct {
	ID        int
	Name      string
	Country   string
	Bio       string
	Birthday  time.Time
	DeathDate time.Time
	Books     []BookListResource
	Images    []MediaResource
}

func NewAuthorResource(model models.Author) *AuthorResource {
	return &AuthorResource{
		ID:        int(model.ID),
		Name:      model.Name,
		Country:   model.Country,
		Bio:       model.Bio,
		Birthday:  model.Birthday,
		DeathDate: model.DeathDate,
		Books:     AsManyBooks(model.Books),
		Images:    AsMediaCollection(model.Images),
	}
}

type AuthorListResource struct {
	ID   int
	Name string
}

func NewAuthorListResource(model models.Author) *AuthorListResource {
	return &AuthorListResource{
		ID:   int(model.ID),
		Name: model.Name,
	}
}
