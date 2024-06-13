package postgres

import (
	"database/sql"
	"dodi/model"
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
    err := c.db.QueryRow("SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses WHERE course_id = $1", id).Scan(
        &course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
    return course, err
}

func (c *CourseRepo) Create(course model.Course) error {
    _, err := c.db.Exec("INSERT INTO courses (title, description) VALUES ($1, $2)", course.Title, course.Description)
    return err
}

func (c *CourseRepo) Update(id string, course model.Course) error {
    _, err := c.db.Exec("UPDATE courses SET title = $1, description = $2 WHERE course_id = $3", course.Title, course.Description, id)
    return err
}

func (c *CourseRepo) Delete(id int) error {
    _, err := c.db.Exec("DELETE FROM courses WHERE course_id = $1", id)
    return err
}
