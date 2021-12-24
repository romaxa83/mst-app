package domains

import "time"

type User struct {
	ID        int       `json:"-" db:"id"`
	Name      string    `json:"name" binding:"required"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password" binding:"required"`
	Status    int8      `json:"status" binding:"required"`
	CreatedAt time.Time `json:"created_at" binding:"required"`
	UpdatedAt time.Time `json:"updated_at" binding:"required"`
}

type Session struct {
	RefreshToken string    `json:"refresh_token" binding:"required"`
	ExpiresAt    time.Time `json:"refresh_token_expires_at" binding:"required"`
}
