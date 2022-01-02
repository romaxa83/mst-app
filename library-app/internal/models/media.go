package models

import "gorm.io/gorm"

type (
	FileStatus int
	FileType   string
	OwnerType  string
)

const (
	MaxUploadSize = 5 << 20 // 5 megabytes
	maxVideoSize  = 2 << 30 // 2 gigabytes
)

var (
	ImageTypes = map[string]interface{}{
		"image/jpeg": nil,
		"image/png":  nil,
	}

	VideoTypes = map[string]interface{}{
		"video/mp4":                 nil,
		"application/octet-stream":  nil,
		"text/plain; charset=utf-8": nil, // for strange files with such content-type
	}

	FileTypes = map[string]interface{}{
		"application/pdf":           nil,
		"application/zip":           nil, // excel
		"text/plain; charset=utf-8": nil, // for strange files with such content-type
	}
)

const (
	ClientUpload FileStatus = iota
	UploadedToStorage
	StorageUploadError
)

const (
	Image FileType = "image"
	Video FileType = "video"
	File  FileType = "file"
)

const (
	AuthorOwner OwnerType = "authors"
)

type Media struct {
	gorm.Model
	OwnerID     int        `json:"owner_id"`
	OwnerType   OwnerType  `json:"owner_type"`
	Type        FileType   `json:"type"`
	ContentType string     `json:"content_type"`
	Name        string     `json:"name"`
	Size        int64      `json:"size"`
	Status      FileStatus `json:"status"`
	URL         string     `json:"url"`
}
