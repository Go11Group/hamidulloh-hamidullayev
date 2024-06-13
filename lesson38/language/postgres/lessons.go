package postgres

import (
	"database/sql"
	"dodi/model"
	_ "github.com/lib/pq"
)

type LessonRepo struct {
	db *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{db: db}
}

func (u *LessonRepo) Get(id string) (model.Lesson, error) {
	var lesson model.Lesson
	err := u.db.QueryRow("SELECT lesson_id , course_id , title, content,created_at, updated_at, deleted_at FROM lessons WHERE lesson_id = $1", id).Scan(
		&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
	return lesson, err
}

func (u *LessonRepo) Create(lesson model.Lesson) error {
	_, err := u.db.Exec("INSERT INTO lessons (course_id, title, content) VALUES ($1, $2, $3)",lesson.CourseID, lesson.Title, lesson.Content)
	return err
}

func (u *LessonRepo) Update(id string, lesson model.Lesson) error {
	_, err := u.db.Exec("UPDATE lessons SET course_id = $1, title = $2, content = $3 WHERE lesson_id = $4",
		lesson.CourseID, lesson.Title, lesson.Content, id)
	return err
}

func (u *LessonRepo) Delete(id string) error {
	_, err := u.db.Exec("DELETE FROM lessons WHERE id = $1", id)
	return err
}