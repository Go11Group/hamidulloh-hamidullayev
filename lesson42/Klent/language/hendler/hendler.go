package hendler

import (
	"github.com/Go11Group/language/postgres"
)

type Handler struct {
	User   *postgres.UserRepo
	Course *postgres.CourseRepo
}
