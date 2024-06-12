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

func (c *CourseRepo) Course_Get(id int) (model.Course, error) {
    var course model.Course
    err := c.db.QueryRow("SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses WHERE course_id = $1", id).Scan(
        &course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &course.DeletedAt)
    return course, err
}

func (c *CourseRepo) Course_Create(course model.Course) error {
    _, err := c.db.Exec("INSERT INTO courses (title, description) VALUES ($1, $2)", course.Title, course.Description)
    return err
}

func (c *CourseRepo) Course_Update(id int, course model.Course) error {
    _, err := c.db.Exec("UPDATE courses SET title = $1, description = $2 WHERE course_id = $3", course.Title, course.Description, id)
    return err
}

func (c *CourseRepo) Course_Delete(id int) error {
    _, err := c.db.Exec("DELETE FROM courses WHERE course_id = $1", id)
    return err
}

func (c *CourseRepo) Courses_GetAll() ([]model.Course, error) {
    rows, err := c.db.Query("SELECT course_id, title, description, created_at, updated_at, deleted_at FROM courses")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    courses := []model.Course{}
    for rows.Next() {
        course := model.Course{}
        var deletedAt sql.NullString
        err = rows.Scan(&course.CourseID, &course.Title, &course.Description, &course.CreatedAt, &course.UpdatedAt, &deletedAt)
        if err != nil {
            return nil, err
        }
        if deletedAt.Valid {
			course.DeletedAt = deletedAt.String
		}
		
		
        courses = append(courses, course)
    }
    return courses, nil
}
