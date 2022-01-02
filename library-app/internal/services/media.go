package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
	"github.com/romaxa83/mst-app/library-app/pkg/logger"
	"github.com/romaxa83/mst-app/library-app/pkg/storage"
	"os"
	"strings"
)

var folders = map[models.FileType]string{
	models.Image: "images",
	models.Video: "videos",
	models.File:  "files",
}

type MediaService struct {
	repo    repositories.Media
	storage storage.Provider
	env     string
}

func NewMediaService(repo repositories.Media, storage storage.Provider, env string) *MediaService {
	return &MediaService{repo, storage, env}
}

func (s MediaService) UploadAndSaveFile(ctx context.Context, input input.UploadMedia) (models.Media, error) {
	defer removeFile(input.Name)

	model, err := s.repo.Create(input)
	if err != nil {
		return models.Media{}, err
	}

	url, err := s.upload(ctx, model)
	if err != nil {
		return models.Media{}, err
	}
	logger.Infof("Upload file - [%s]", url)

	return s.repo.SetUrl(model, url)
}

func (s *MediaService) upload(ctx context.Context, file models.Media) (string, error) {
	f, err := os.Open(file.Name)
	if err != nil {
		return "", err
	}

	info, _ := f.Stat()
	logger.Infof("file info: %+v", info)

	defer f.Close()

	return s.storage.Upload(ctx, storage.UploadInput{
		File:        f,
		Size:        file.Size,
		ContentType: file.ContentType,
		Name:        s.generateFilename(file),
	})
}

func (s *MediaService) generateFilename(file models.Media) string {
	filename := fmt.Sprintf("%s.%s", uuid.New().String(), getFileExtension(file.Name))
	folder := folders[file.Type]

	fileNameParts := strings.Split(file.Name, "-")

	return fmt.Sprintf("%s/%s/%s/%s/%s", s.env, file.OwnerType, fileNameParts[0], folder, filename)
}

func getFileExtension(filename string) string {
	parts := strings.Split(filename, ".")

	return parts[len(parts)-1]
}

func removeFile(filename string) {
	if err := os.Remove(filename); err != nil {
		logger.Error("removeFile(): ", err)
	}
}
