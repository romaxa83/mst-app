package domains

import "time"

type (
	FileStatus int
	FileType   string
)

const (
	ClientUploadInProgress FileStatus = iota
	UploadedByClient
	ClientUploadError
	StorageUploadInProgress
	UploadedToStorage
	StorageUploadError
)

const (
	Image FileType = "image"
	Video FileType = "video"
	Other FileType = "other"
)

type File struct {
	ID              int        `json:"-" db:"id"`
	Type            FileType   `json:"type"`
	ContentType     string     `json:"contentType" db:"content_type"`
	Name            string     `json:"name"`
	Size            int64      `json:"size"`
	Status          FileStatus `json:"status"`
	UploadStartedAt time.Time  `json:"uploadStartedAt" bson:"upload_started_at"`
	URL             string     `json:"url"`
}
