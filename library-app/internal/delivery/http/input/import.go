package input

import "github.com/romaxa83/mst-app/library-app/internal/models"

type CreateImport struct {
	Entity      models.ImportEntity      `json:"entity"`
	ContentType models.ImportContentType `json:"type"`
	Status      models.ImportStatus      `json:"status"`
	FilePath    string                   `json:"file_path"`
	Message     string                   `json:"message"`
}
