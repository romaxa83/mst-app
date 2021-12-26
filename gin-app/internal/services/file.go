package services

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"github.com/romaxa83/mst-app/gin-app/internal/repositories"
	"github.com/romaxa83/mst-app/gin-app/pkg/logger"
	"github.com/romaxa83/mst-app/gin-app/pkg/storage"
	"os"
	"strings"
	"time"
)

const (
	_workersCount   = 2
	_workerInterval = time.Second * 10
)

var folders = map[domains.FileType]string{
	domains.Image: "images",
	domains.Video: "videos",
	domains.Other: "other",
}

type FilesService struct {
	repo    repositories.Files
	storage storage.Provider
	env     string
}

func NewFilesService(repo repositories.Files, storage storage.Provider, env string) *FilesService {
	return &FilesService{repo: repo, storage: storage, env: env}
}

// сохраняем данные по файлу в бд

func (s *FilesService) Save(ctx context.Context, file domains.File) (domains.File, error) {
	return s.repo.Create(ctx, file)
}

func (s *FilesService) UpdateStatus(ctx context.Context, fileName string, status domains.FileStatus) error {
	return s.repo.UpdateStatus(ctx, fileName, status)
}

func (s *FilesService) GetByID(ctx context.Context, id string) (domains.File, error) {
	return s.repo.GetByID(ctx, id)
}

// сохраняем данные по файлу в бд и грузим картинку на сторонний сервис

func (s *FilesService) UploadAndSaveFile(ctx context.Context, file domains.File) (string, error) {
	defer removeFile(file.Name)

	file.UploadStartedAt = time.Now()

	fileSave, err := s.Save(ctx, file)
	if err != nil {
		return "", err
	}

	url, err := s.upload(ctx, file)
	if err != nil {
		return "", err
	}
	logger.Infof("Upload file - [%s]", url)

	s.repo.UpdateStatusAndSetURL(ctx, fmt.Sprintf("%v", fileSave.ID), url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (s *FilesService) InitStorageUploaderWorkers(ctx context.Context) {
	for i := 0; i < _workersCount; i++ {
		go s.processUploadToStorage(ctx)
		logger.Warn("GO-GO-UPLOADER-WORKER")
	}
}

func (s *FilesService) processUploadToStorage(ctx context.Context) {
	for {
		if err := s.uploadToStorage(ctx); err != nil {
			logger.Error("uploadToStorage(): ", err)
			logger.Errorf("%+v", err)
		}

		time.Sleep(_workerInterval)
	}
}

func (s *FilesService) uploadToStorage(ctx context.Context) error {
	file, err := s.repo.GetForUploading(ctx)
	if err != nil {
		return err
	}
	defer removeFile(file.Name)
	id := fmt.Sprintf("%s", file.ID)

	logger.Infof("processing file %s", file.Name)

	url, err := s.upload(ctx, file)
	if err != nil {
		if err := s.repo.UpdateStatus(ctx, id, domains.StorageUploadError); err != nil {
			return err
		}

		return err
	}

	logger.Infof("file %s processed successfully", file.Name)

	if err := s.repo.UpdateStatusAndSetURL(ctx, id, url); err != nil {
		return err
	}

	return nil
}

func (s *FilesService) upload(ctx context.Context, file domains.File) (string, error) {
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
		Type:        file.Type,
		Name:        s.generateFilename(file),
	})
}

func (s *FilesService) generateFilename(file domains.File) string {
	filename := fmt.Sprintf("%s.%s", uuid.New().String(), getFileExtension(file.Name))

	return fmt.Sprintf("%s/%s", file.Type, filename)
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
