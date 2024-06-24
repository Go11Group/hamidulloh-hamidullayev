package handler

import (
	"database/sql"

	postgres "github.com/Go11Group/lesson43/metro_service/storage"
)

type handler struct {
	Card *postgres.CardRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		Card: postgres.NewCardRepo(db),
	}
}
