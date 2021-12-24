package storage

import (
	"context"
	"fmt"
)

type FileDiscStorage struct {
	bucket   string
	endpoint string
}

func NewFileStorage(bucket, endpoint string) *FileDiscStorage {
	return &FileDiscStorage{
		bucket:   bucket,
		endpoint: endpoint,
	}
}

func (fs *FileDiscStorage) Upload(ctx context.Context, input UploadInput) (string, error) {

	return fs.generateFileURL(input.Name), nil
}

// DigitalOcean Spaces URL format.
func (fs *FileDiscStorage) generateFileURL(filename string) string {
	return fmt.Sprintf("https://%s.%s/%s", fs.bucket, fs.endpoint, filename)
}
