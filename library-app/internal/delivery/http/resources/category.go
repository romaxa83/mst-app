package resources

import "github.com/romaxa83/mst-app/library-app/internal/models"

type CategoryResource struct {
	ID     int
	Title  string
	Desc   string
	Active bool
	Sort   int
	Books  []BookListResource
}

func NewCategoryResource(model models.Category) *CategoryResource {
	return &CategoryResource{
		ID:     int(model.ID),
		Title:  model.Title,
		Desc:   model.Desc,
		Active: model.Active,
		Sort:   model.Sort,
		Books:  AsManyBooks(model.Books),
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

func AsManyCategories(models []models.Category) []CategoryListResource {
	var list []CategoryListResource
	for _, b := range models {
		list = append(list, *NewCategoryListResource(b))
	}
	return list
}
