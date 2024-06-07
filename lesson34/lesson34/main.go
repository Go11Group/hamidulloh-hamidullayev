package main

import (
	"github.com/Go11Group/at_lesson/lesson34/handler"
	"github.com/Go11Group/at_lesson/lesson34/postgres"
)

func main() {
	db, err := postgres.User_ConnectDB()
	if err != nil {
		panic(err)
	}
	// usere := postgres.User_NewBookRepo(db)
	product := postgres.Product_NewBookRepo(db)

	server := handler.NewHandler(handler.Handler{Productr:  product})

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}