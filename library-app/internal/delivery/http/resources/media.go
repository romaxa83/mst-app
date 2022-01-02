package resources

import (
	"github.com/romaxa83/mst-app/library-app/internal/models"
)

type MediaResource struct {
	ID          int
	Name        string
	ContentType string
	Size        int64
	URL         string
}

func NewMediaResource(model models.Media) *MediaResource {
	return &MediaResource{
		ID:          int(model.ID),
		Name:        model.Name,
		ContentType: model.ContentType,
		Size:        model.Size,
		URL:         model.URL,
	}
}

func AsMediaCollection(models []models.Media) []MediaResource {
	var list []MediaResource
	for _, b := range models {
		list = append(list, *NewMediaResource(b))
	}
	return list
}
