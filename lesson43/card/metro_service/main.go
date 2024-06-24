package main

import (
	"github.com/Go11Group/lesson43/metro_service/api"
	postgres "github.com/Go11Group/lesson43/metro_service/storage"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	server := api.Routes(db)
	err = server.ListenAndServe() 
	if err != nil {
		panic(err)
	}
}