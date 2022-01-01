package resources

import (
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"time"
)

type BookResource struct {
	ID          int
	Title       string
	Desc        string
	Pages       int
	Qty         int
	Sort        int
	Active      bool
	PublishedAt time.Time
	Author      *AuthorResource
}

func NewBookResource(model models.Book) *BookResource {
	return &BookResource{
		int(model.ID),
		model.Title,
		model.Desc,
		model.Pages,
		model.Qty,
		model.Sort,
		model.Active,
		model.PublishedAt,
		NewAuthorResource(model.Author),
	}
}

type BookListResource struct {
	ID          int
	Title       string
	Sort        int
	Active      bool
	PublishedAt time.Time
}

func NewBookListResource(model models.Book) *BookListResource {
	return &BookListResource{
		int(model.ID),
		model.Title,
		model.Sort,
		model.Active,
		model.PublishedAt,
	}
}

func AsManyBooks(models []models.Book) []BookListResource {
	var list []BookListResource
	for _, b := range models {
		list = append(list, *NewBookListResource(b))
	}
	return list
}
