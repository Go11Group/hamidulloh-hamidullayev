package model

import (
	"time"

	"github.com/google/uuid"
)

type Enrollment struct {
	EnrollmentID   uuid.UUID `json:"enrollment_id"`
	UserID         uuid.UUID `json:"user_id"`
	CourseID       uuid.UUID `json:"course_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      int       `json:"deleted_at"`
}
