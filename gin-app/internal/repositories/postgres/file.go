package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
)

type FilesRepo struct {
	db *sqlx.DB
}

func NewFilesRepo(db *sqlx.DB) *FilesRepo {
	return &FilesRepo{
		db: db,
	}
}

func (r *FilesRepo) Create(ctx context.Context, file domains.File) (domains.File, error) {
	var model domains.File
	query := fmt.Sprintf(`INSERT INTO %s
									(type, content_type, name, size, upload_started_at)
								values ($1, $2, $3, $4, $5) RETURNING id, name, type, content_type, size, status`,
		mediaTable)

	row := r.db.QueryRow(
		query,
		file.Type,
		file.ContentType,
		file.Name,
		file.Size,
		file.UploadStartedAt,
	)

	if err := row.Scan(
		&model.ID,
		&model.Name,
		&model.Type,
		&model.ContentType,
		&model.Size,
		&model.Status,
	); err != nil {
		return domains.File{}, err
	}

	return model, nil
}

func (r *FilesRepo) UpdateStatus(ctx context.Context, id string, status domains.FileStatus) error {
	query := fmt.Sprintf(`UPDATE %s m
								SET status = $2
								WHERE m.id = %s`,
		mediaTable, id)

	_, err := r.db.Exec(query, status)

	return err
}

func (r *FilesRepo) GetForUploading(ctx context.Context) (domains.File, error) {
	var file domains.File

	query := fmt.Sprintf(`SELECT id, type FROM %s 
						WHERE status = $1 LIMIT 1`,
		mediaTable)
	err := r.db.Get(&file, query, domains.StorageUploadInProgress)

	return file, err
}

func (r *FilesRepo) UpdateStatusAndSetURL(ctx context.Context, id, url string) error {
	query := fmt.Sprintf(`UPDATE %s m
								SET url = $1, status = $2
								WHERE m.id = %s`,
		mediaTable, id)
	_, err := r.db.Exec(query, url, domains.UploadedToStorage)

	return err
}

func (r *FilesRepo) GetByID(ctx context.Context, id string) (domains.File, error) {
	var file domains.File

	query := fmt.Sprintf(`SELECT * FROM %s 
						WHERE id = $1 LIMIT 1`,
		mediaTable)

	err := r.db.Get(&file, query, id)

	return file, err
}
