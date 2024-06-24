package handler

import (
	"database/sql"

	postgres "github.com/Go11Group/lesson43/metro_service/storage"
)

type handler struct {
	User *postgres.UserRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		User: postgres.NewUserRepo(db),
	}
}
