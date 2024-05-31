package modul

import (
	_"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName  string
	LastName   string
	Email      string
	Password   string
	Age        int
	Fiald      string
	Gander     string
	IsEmployee bool
}
