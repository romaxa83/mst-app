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

func (r *FilesRepo) Create(ctx context.Context, file domains.File) (int, error) {

	var id int
	query := fmt.Sprintf(`INSERT INTO %s
									(type, content_type, name, size, upload_started_at)
								values ($1, $2, $3, $4, $5) RETURNING id`,
		mediaTable)

	row := r.db.QueryRow(
		query,
		file.Type,
		file.ContentType,
		file.Name,
		file.Size,
		file.UploadStartedAt,
	)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

//func (r *FilesRepo) UpdateStatus(ctx context.Context, fileName string, status domain.FileStatus) error {
//	_, err := r.db.UpdateOne(ctx, bson.M{"name": fileName}, bson.M{"$set": bson.M{"status": status}})
//
//	return err
//}
//
//func (r *FilesRepo) GetForUploading(ctx context.Context) (domain.File, error) {
//	var file domain.File
//
//	res := r.db.FindOneAndUpdate(ctx, bson.M{"status": domain.UploadedByClient}, bson.M{"$set": bson.M{"status": domain.StorageUploadInProgress}})
//	err := res.Decode(&file)
//
//	return file, err
//}
//
//func (r *FilesRepo) UpdateStatusAndSetURL(ctx context.Context, id primitive.ObjectID, url string) error {
//	_, err := r.db.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{"url": url, "status": domain.UploadedToStorage}})
//
//	return err
//}
//
//func (r *FilesRepo) GetByID(ctx context.Context, id, schoolId primitive.ObjectID) (domain.File, error) {
//	var file domain.File
//
//	res := r.db.FindOne(ctx, bson.M{"_id": id, "schoolId": schoolId})
//	err := res.Decode(&file)
//
//	return file, err
//}
