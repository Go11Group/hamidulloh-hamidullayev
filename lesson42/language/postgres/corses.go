package postgres

import (
	"database/sql"
	"time"

	"github.com/Go11Group/language/model"

	_ "github.com/lib/pq"
)

// CourseRepo handles CRUD operations for courses in the database.
type CourseRepo struct {
	db *sql.DB
}

// NewCourseRepo creates a new CourseRepo instance.
func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

// Get retrieves a course by its ID.
func (c *CourseRepo) Get(id string) (model.Course, error) {
	var course model.Course
	err := c.db.QueryRow("SELECT course_id, title, description, created_at, update_at FROM course WHERE course_id = $1 and deleted_at = 0", id).Scan(
		&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt)
	return course, err
}

// Create adds a new course to the database.
func (c *CourseRepo) Create(course model.Course) error {
	_, err := c.db.Exec("INSERT INTO course (title, description, deleted_at, update_at) VALUES ($1, $2, $3, $4)", course.Title, course.Description, 0, time.Now())
	return err
}

// Update modifies an existing course in the database.
func (c *CourseRepo) Update(id string, course model.Course) error {
	_, err := c.db.Exec("UPDATE course SET title = $1, description = $2, update_at = $4 WHERE course_id = $3 and deleted_at = 0", course.Title, course.Description, id, time.Now())
	return err
}

// Delete marks a course as deleted in the database.
func (c *CourseRepo) Delete(id string) error {
	_, err := c.db.Exec("update course set deleted_at = date_part('epoch', current_timestamp)::INT where course_id = $1 and deleted_at = 0", id)
	return err
}
