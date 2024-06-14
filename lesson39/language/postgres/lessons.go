package postgres

import (
	"database/sql"
	"time"

	"github.com/Go11Group/language/model"

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
	err := u.db.QueryRow("SELECT lesson_id , course_id , title, content, created_at, update_at FROM lessons WHERE lesson_id = $1 and deleted_at = 0", id).Scan(
		&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt)
	return lesson, err
}

func (u *LessonRepo) Create(lesson model.Lesson) error {
	_, err := u.db.Exec("INSERT INTO lessons (course_id, title, content, deleted_at, update_at) VALUES ($1, $2, $3, $4, $5)", lesson.CourseID, lesson.Title, lesson.Content, 0, time.Now())
	return err
}

func (u *LessonRepo) Update(id string, lesson model.Lesson) error {
	_, err := u.db.Exec("UPDATE lessons SET course_id = $1, title = $2, content = $3, update_at = $5 WHERE lesson_id = $4 and deleted_at = 0",
		lesson.CourseID, lesson.Title, lesson.Content, id, time.Now())
	return err
}

func (u *LessonRepo) Delete(id string) error {
	_, err := u.db.Exec("update lessons set deleted_at = date_part('epoch', current_timestamp)::INT where lesson_id = $1 and deleted_at = 0", id)
	return err
}
func (u *LessonRepo) GetAll(a string, arr []interface{}) ([]model.Lesson, error) {
	rows, err := u.db.Query(a, arr...)
	if err != nil {
		return nil, err
	}
	lessons := []model.Lesson{}
	for rows.Next() {
		lesson := model.Lesson{}
		err = rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content,
		&lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

func (h *LessonRepo) LessonsByCourse(courseID string) ([]model.Lesson, error) {
	query := `SELECT lesson_id, course_id, title, content, created_at, update_at, deleted_at 
              FROM lessons 
              WHERE course_id = $1 AND deleted_at = 0`

	rows, err := h.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []model.Lesson
	for rows.Next() {
		var lesson model.Lesson
		err := rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content, &lesson.CreatedAt, &lesson.UpdatedAt, &lesson.DeletedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}
