package funcgo

import (
	"Modul/modul"
	"fmt"

	// "fmt"

	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbUser struct {
	DB *gorm.DB
}

func (db *dbUser) Create1(user modul.User) {
	db.DB.Create(&user)
}
func (db *dbUser) Delete1(userID uint) {
	db.DB.Delete(&modul.User{}, userID)
}

func (db *dbUser) Update1(user modul.User, id string) error {
	result := db.DB.Model(&modul.User{}).Where("id = ?", id).Updates(user)
	return result.Error
}

func (db *dbUser) FindByID(id string, a int) (*modul.User, error) {
	switch a {
	case 1:
		writer := &modul.User{}
		if err := db.DB.First(writer, id).Error; err != nil {
			return nil, err
		}
		return writer, nil
	case 2:
		writer := &modul.User{}
		if err := db.DB.Where(fmt.Sprintf("%s = ?", "first_name"), id).First(writer).Error; err != nil {
			return nil, err
		}
		return writer, nil
	case 3:
		writer := &modul.User{}
		if err := db.DB.Where(fmt.Sprintf("%s = ?", "email"), id).First(writer).Error; err != nil {
			return nil, err
		}
		return writer, nil
	case 4:
		writer := &modul.User{}
		if err := db.DB.Where(fmt.Sprintf("%s = ?", "last_name"), id).First(writer).Error; err != nil {
			return nil, err
		}
		return writer, nil
	}

	return nil,fmt.Errorf("sizqidirgan narsa chiqmadi")
}


func NewDBUser() (*dbUser, error) {
	db, err := gorm.Open(postgres.Open("postgres://postgres:dodi@localhost:5432/dodi?sslmode=disable"))
	if err != nil {
		return nil, err
	}
	return &dbUser{DB: db}, nil
}
