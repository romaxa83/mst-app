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
