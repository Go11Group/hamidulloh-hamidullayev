package hendler

import (
	"dodi/postgres"
)

type Handler struct {
	User       *postgres.UserRepo
	UserFilter *postgres.UserFilterRepo
	Course     *postgres.CourseRepo
	Enrollment *postgres.EnrollmentRepo
	Lesson     *postgres.LessonRepo
}
