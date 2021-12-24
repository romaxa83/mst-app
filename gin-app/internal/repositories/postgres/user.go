package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/romaxa83/mst-app/gin-app/internal/domains"
	"time"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user domains.User) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s 
									(name, email, phone, password, status, created_at, updated_at) 
								values ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		usersTable)

	row := r.db.QueryRow(query, user.Name, user.Email, user.Phone, user.Password, user.Status, user.CreatedAt, user.UpdatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepo) GetByCredentials(ctx context.Context, email, password string) (domains.User, error) {
	var user domains.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password=$2", usersTable)

	err := r.db.Get(&user, query, email, password)

	return user, err
}

func (r *UserRepo) GetByRefreshToken(ctx context.Context, refreshToken string) (domains.User, error) {
	var user domains.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE refresh_token=$1 AND refresh_token_expires_at>$2 ", usersTable)

	err := r.db.Get(&user, query, refreshToken, time.Now())

	return user, err
}

func (r *UserRepo) SetSession(ctx context.Context, userId int, session domains.Session) error {
	query := fmt.Sprintf(`UPDATE %s SET refresh_token=$1, refresh_token_expires_at=$2
								WHERE id=$3`, usersTable)

	_, err := r.db.Exec(query, session.RefreshToken, session.ExpiresAt, userId)

	return err
}
