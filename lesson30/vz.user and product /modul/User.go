package modul

import (
	_ "gorm.io/driver/postgres"
)

type User struct {
	Id       int
	UserName string
	Email    string
	Password string
}

type Products struct {
	Id             uint
	Name           string
	Description    string
	Price          float32 
	Stock_quantity int    
}
