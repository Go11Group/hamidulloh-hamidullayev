package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Go11Group/language/model"

	_ "github.com/lib/pq"
)

type CourseRepo struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

func (c *CourseRepo) Get(id string) (model.Course, error) {
	var course model.Course
	err := c.db.QueryRow("SELECT course_id, title, description, created_at, update_at FROM course WHERE course_id = $1 and deleted_at = 0", id).Scan(
		&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt)
	fmt.Println(course)
	return course, err
}

func (c *CourseRepo) Create(course model.Course) error {
	_, err := c.db.Exec("INSERT INTO course (title, description, deleted_at, update_at) VALUES ($1, $2, $3, $4)", course.Title, course.Description, 0, time.Now())
	return err
}

func (c *CourseRepo) Update(id string, course model.Course) error {
	_, err := c.db.Exec("UPDATE courses SET title = $1, description = $2, update_at = $4 WHERE course_id = $3 and deleted_at = 0", course.Title, course.Description, id, time.Now())
	return err
}

func (c *CourseRepo) Delete(id int) error {
	_, err := c.db.Exec("update courses set deleted_at = date_part('epoch', current_timestamp)::INT where id = $1 and deleted_at = 0", id)
	return err
}
func (u *CourseRepo) GetAll(a string, arr []interface{}) ([]model.Course, error) {
	rows, err := u.db.Query(a, arr...)
	if err != nil {
		return nil, err
	}
	courses := []model.Course{}
	for rows.Next() {
		course := model.Course{}
		err = rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt,
			&course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (u *CourseRepo) CoursesByUser(userID string) ([]model.Course, error) {
	query := `
	SELECT c.course_id, c.title, c.description, c.created_at, c.update_at, c.deleted_at
	FROM Course c
	INNER JOIN Enrollment e ON c.course_id = e.course_id
	WHERE e.user_id = $1 AND c.deleted_at = 0
	`
	rows, err := u.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []model.Course
	for rows.Next() {
		var course model.Course
		err := rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}
