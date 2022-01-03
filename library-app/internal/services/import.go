package services

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/romaxa83/mst-app/library-app/internal/delivery/http/input"
	"github.com/romaxa83/mst-app/library-app/internal/models"
	"github.com/romaxa83/mst-app/library-app/internal/repositories"
	"github.com/romaxa83/mst-app/library-app/pkg/file"
	"os"
	"time"
)

type ImportService struct {
	repo       repositories.Import
	repoAuthor repositories.Author
}

func NewImportService(repo repositories.Import, repoAuthor repositories.Author) *ImportService {
	return &ImportService{repo, repoAuthor}
}

func (s *ImportService) Create(input input.CreateImport) (models.Import, error) {
	return s.repo.Create(input)
}

func (s *ImportService) Parse(model models.Import) error {

	if !file.Exists(model.FilePath) {
		return errors.New(fmt.Sprintf("file by path [%s] not exist", model.FilePath))
	}

	f, err := os.Open(model.FilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}

	layout := "2006-01-02"

	for _, row := range rows {
		if m, _ := s.repoAuthor.GetOneByName(row[0]); m.ID == 0 {
			bdate, _ := time.Parse(layout, row[2])
			ddate, _ := time.Parse(layout, row[3])
			a := input.CreateAuthor{
				Name:      row[0],
				Country:   row[1],
				Birthday:  bdate,
				DeathDate: ddate,
			}
			s.repoAuthor.Create(a)
		}
	}

	return nil
}
