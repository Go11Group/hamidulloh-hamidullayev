package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Go11Group/language/model"
	_ "github.com/lib/pq"
)

type EnrollmentRepo struct {
	db *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{db: db}
}

func (er *EnrollmentRepo) Get(id string) (model.Enrollment, error) {
	var enrollment model.Enrollment
	err := er.db.QueryRow("SELECT enrollment_id, user_id, course_id, enrollment_data, created_at, update_at FROM enrollment WHERE enrollment_id = $1 and deleted_at = 0", id).Scan(
		&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate, &enrollment.CreatedAt, &enrollment.UpdatedAt )
	return enrollment, err
}

func (er *EnrollmentRepo) Create(enrollment model.Enrollment) error {
	fmt.Println(enrollment)
	_, err := er.db.Exec("INSERT INTO enrollment (user_id, course_id, enrollment_data, update_at, deleted_at) VALUES ($1, $2, $3, $4, $5)", enrollment.UserID, enrollment.CourseID, time.Now(), time.Now(), 0)
	return err
}

func (er *EnrollmentRepo) Update(id string, enrollment model.Enrollment) error {
	_, err := er.db.Exec("UPDATE enrollment SET user_id = $1, course_id = $2, update_at = $3 WHERE enrollment_id = $4 and deleted_at = 0", enrollment.UserID, enrollment.CourseID, time.Now(),id)
	return err
}

func (er *EnrollmentRepo) Delete(id string) error {
	_, err := er.db.Exec("UPDATE enrollment SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE enrollment_id = $1 AND deleted_at = 0", id)
	return err
}

func (u *EnrollmentRepo) EnrolledUsersByCourse(courseID string) ([]model.User, error) {
	query := `SELECT u.user_id, u.name, u.email, u.birthday, u.password, u.created_at, u.update_at, u.deleted_at 
              FROM users u
              JOIN enrollment e ON u.user_id = e.user_id
              WHERE e.course_id = $1 AND u.deleted_at = 0`

	rows, err := u.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
