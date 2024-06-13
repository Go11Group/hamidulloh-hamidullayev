package postgres

import (
	"database/sql"
	"dodi/model"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type EnrollmentRepo struct {
	db *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{db: db}
}

func (er *EnrollmentRepo) Get(id uuid.UUID) (model.Enrollment, error) {
	var enrollment model.Enrollment
	err := er.db.QueryRow("SELECT enrollment_id, user_id, course_id, enrollment_date, created_at, updated_at, deleted_at FROM enrollments WHERE enrollment_id = $1", id).Scan(
		&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt, &enrollment.DeletedAt)
	return enrollment, err
}

func (er *EnrollmentRepo) Create(enrollment model.Enrollment) error {
	_, err := er.db.Exec("INSERT INTO enrollments (user_id, course_id, enrollment_date) VALUES ($1, $2, $3)", enrollment.UserID, enrollment.CourseID, time.Now())
	return err
}

func (er *EnrollmentRepo) Update(id uuid.UUID, enrollment model.Enrollment) error {
	_, err := er.db.Exec("UPDATE enrollments SET user_id = $1, course_id = $2, enrollment_date = $3 WHERE enrollment_id = $4", enrollment.UserID, enrollment.CourseID, enrollment.EnrollmentDate, id)
	return err
}

func (er *EnrollmentRepo) Delete(id uuid.UUID) error {
	_, err := er.db.Exec("DELETE FROM enrollments WHERE enrollment_id = $1", id)
	return err
}
