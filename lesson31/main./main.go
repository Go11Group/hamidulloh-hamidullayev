package main

import (
	"database/sql"
	"fmt"

	// "math/rand"

	"github.com/go-faker/faker/v3"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "postgres://postgres:dodi@localhost:5432/home?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for i := 0; i < 100; i++ {
		_, err = db.Exec("insert into dodi(name,surname,phonenumber,email,password) values ($1,$2,$3,$4,$5)", faker.FirstName(), faker.LastName(), faker.PhoneNumber, faker.Email(), faker.Password())
		if err != nil {
			panic(err)
		}
		if i%100 == 0 {
			fmt.Println(i)
		}
	}
}