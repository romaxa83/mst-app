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
	Author      *AuthorListResource
	Categories  []CategoryListResource
}

func NewBookResource(model models.Book) *BookResource {
	return &BookResource{
		ID:          int(model.ID),
		Title:       model.Title,
		Desc:        model.Desc,
		Pages:       model.Pages,
		Qty:         model.Qty,
		Sort:        model.Sort,
		Active:      model.Active,
		PublishedAt: model.PublishedAt,
		Author:      NewAuthorListResource(model.Author),
		Categories:  AsManyCategories(model.Categories),
	}
}

type BookListResource struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func NewBookListResource(model models.Book) *BookListResource {
	return &BookListResource{
		int(model.ID),
		model.Title,
	}
}

func AsManyBooks(models []models.Book) []BookListResource {
	var list []BookListResource
	for _, b := range models {
		list = append(list, *NewBookListResource(b))
	}
	return list
}
