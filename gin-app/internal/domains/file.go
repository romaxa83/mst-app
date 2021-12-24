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
	ContentType     string     `json:"contentType" bson:"contentType"`
	Name            string     `json:"name" bson:"name"`
	Size            int64      `json:"size" bson:"size"`
	Status          FileStatus `json:"status" bson:"status,omitempty"`
	UploadStartedAt time.Time  `json:"uploadStartedAt" bson:"uploadStartedAt"`
	URL             string     `json:"url" bson:"url,omitempty"`
}
