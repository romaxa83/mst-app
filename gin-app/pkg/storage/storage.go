package storage

import (
	"context"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"io"
)

type UploadInput struct {
	File        io.Reader
	Name        string
	Size        int64
	ContentType string
	Type        domains.FileType
}

type Provider interface {
	Upload(ctx context.Context, input UploadInput) (string, error)
}
