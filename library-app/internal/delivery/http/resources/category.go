package resources

import "github.com/romaxa83/mst-app/library-app/internal/models"

type CategoryResource struct {
	ID     int
	Title  string
	Desc   string
	Active bool
	Sort   int
}

func NewCategoryResource(model models.Category) *CategoryResource {
	return &CategoryResource{
		ID:     int(model.ID),
		Title:  model.Title,
		Desc:   model.Desc,
		Active: model.Active,
		Sort:   model.Sort,
	}
}

type CategoryListResource struct {
	ID    int
	Title string
}

func NewCategoryListResource(model models.Category) *CategoryListResource {
	return &CategoryListResource{
		ID:    int(model.ID),
		Title: model.Title,
	}
}
