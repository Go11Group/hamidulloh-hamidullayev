package hendler

import (
	"dodi/postgres"
)

type Handler struct {
	User          *postgres.UserRepo
	Course       *postgres.CourseRepo
}