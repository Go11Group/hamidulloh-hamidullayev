package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Birthday  string `json:"birthday,omitempty"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt int       `json:"deleted_at"`
}

type UserFilter struct {
	Name      *string    `json:"name"`
	Email     *string    `json:"email"`
	Birthday  *time.Time `json:"birthday"`
	Password  *string    `json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Limit     *int       `json:"limit"`
	Offset    *int       `json:"offset"`
}