package input

import "github.com/romaxa83/mst-app/library-app/pkg/db"

type CreateCategory struct {
	Title string `json:"title" binding:"required,max=256"`
	Desc  string `json:"desc"`
}

// все поля указатели чтобы при обновлении понять какие имеют значения, и не заполнены нулевыми значениями

type UpdateCategory struct {
	Title  *string `json:"title" binding:"max=256"`
	Desc   *string `json:"desc"`
	Active *bool   `json:"active"`
	Sort   *int    `json:"sort"`
}

type CategoryFilterQuery struct {
	Active *bool `form:"active"`
	Sort   *int  `form:"sort"`
	Id     *int  `form:"id"`
}

type GetCategoryQuery struct {
	Pagination db.Pagination
	Search     db.SearchQuery
	CategoryFilterQuery
}
