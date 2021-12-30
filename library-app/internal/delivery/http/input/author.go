package input

import (
	"github.com/romaxa83/mst-app/library-app/pkg/db"
	"time"
)

type CreateAuthor struct {
	Name      string    `json:"name" binding:"required,max=256"`
	Country   string    `json:"country" binding:"required,max=256"`
	Bio       string    `json:"bio" binding:"required"`
	Birthday  time.Time `json:"birthday" binding:"required" time_format:"2006-01-02" time_utc:"1"`
	DeathDate time.Time `json:"death_date" time_format:"2006-01-02" time_utc:"1"`
}

// все поля указатели чтобы при обновлении понять какие имеют значения, и не заполнены нулевыми значениями

type UpdateAuthor struct {
	Name      *string    `json:"name" binding:"max=256"`
	Country   *string    `json:"country" binding:"max=256"`
	Bio       *string    `json:"bio"`
	Birthday  *time.Time `json:"birthday" time_format:"2006-01-02" time_utc:"1"`
	DeathDate *time.Time `json:"death_date" time_format:"2006-01-02" time_utc:"1"`
}

type AuthorFilterQuery struct {
	Id *int `form:"id"`
}

type GetAuthorQuery struct {
	Pagination db.Pagination
	Search     db.SearchQuery
	AuthorFilterQuery
}
