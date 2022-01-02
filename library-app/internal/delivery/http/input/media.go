package input

import (
	"github.com/romaxa83/mst-app/library-app/internal/models"
)

type UploadMedia struct {
	OwnerID     int               `json:"owner_id"`
	OwnerType   models.OwnerType  `json:"owner_type"`
	Type        models.FileType   `json:"type"`
	ContentType string            `json:"contentType" db:"content_type"`
	Name        string            `json:"name"`
	Size        int64             `json:"size"`
	Status      models.FileStatus `json:"status"`
}
