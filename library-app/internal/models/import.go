package models

import "gorm.io/gorm"

type (
	ImportEntity      string
	ImportContentType string
	ImportStatus      int
)

const (
	AuthorEntityImport ImportEntity = "author"
)

const (
	CSVImport ImportContentType = "text/csv"
)

const (
	ImportUpload ImportStatus = iota
	ImportProcess
	ImportDone
	ImportUploadStorage
	ImportError
)

var (
	ImportTypes = map[string]interface{}{
		string(CSVImport): nil,
	}
)

type Import struct {
	gorm.Model
	Entity      ImportEntity      `json:"entity" gorm:"size:50"`
	ContentType ImportContentType `json:"content_type" gorm:"size:250"`
	Status      ImportStatus      `json:"status" gorm:"default:0"`
	FilePath    string            `json:"file_path"`
	Message     string            `json:"message"`
}
