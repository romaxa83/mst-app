package dto

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreateAuthorDto struct {
	ID       uuid.UUID `json:"id" validate:"required"`
	Name     string    `json:"name" validate:"required,gte=0,lte=255"`
	Bio      string    `json:"bio" validate:"required,gte=0,lte=5000"`
	Birthday time.Time `json:"price" validate:"required" time_format:"2006-01-02" time_utc:"1"`
}

type CreateAuthorResponseDto struct {
	ID uuid.UUID `json:"id" validate:"required"`
}
