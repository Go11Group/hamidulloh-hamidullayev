package model

import (
	"time"
	"github.com/google/uuid"
)

type Course struct {
	CourseID   uuid.UUID  `json:"course_id"`
	Title      string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  string `json:"deleted_at,omitempty"`
}