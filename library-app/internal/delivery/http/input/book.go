package input

import (
	"github.com/romaxa83/mst-app/library-app/pkg/db"
	"time"
)

type CreateBook struct {
	Title       string    `json:"title" binding:"required,max=256"`
	Desc        string    `json:"desc"`
	PublishedAt time.Time `json:"published_at" binding:"required" time_format:"2006-01-02" time_utc:"1"`
	Pages       int       `json:"pages" binding:"required"`
	Qty         int       `json:"qty"`
	AuthorID    int       `json:"author_id" binding:"required"`
}

// все поля указатели чтобы при обновлении понять какие имеют значения, и не заполнены нулевыми значениями

type UpdateBook struct {
	Title       *string    `json:"title" binding:"max=256"`
	Desc        *string    `json:"desc"`
	PublishedAt *time.Time `json:"published_at" time_format:"2006-01-02" time_utc:"1"`
	Pages       *int       `json:"pages"`
	Qty         *int       `json:"qty"`
	AuthorID    *int       `json:"author_id"`
	Sort        *int       `json:"sort"`
	Active      *bool      `json:"active"`
}

type BookFilterQuery struct {
	Active *bool `form:"active"`
	Sort   *int  `form:"sort"`
	Id     *int  `form:"id"`
}

type GetBookQuery struct {
	Pagination db.Pagination
	Search     db.SearchQuery
	BookFilterQuery
}
