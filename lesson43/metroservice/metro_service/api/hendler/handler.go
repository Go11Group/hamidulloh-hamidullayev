package handler

import (
	"database/sql"

	postgres "github.com/Go11Group/lesson43/metro_service/storage"
)

type handler struct {
	Station *postgres.StationRepo
	Terminal *postgres.TerminalRepo
	Transaction *postgres.TransactionRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		Station: postgres.NewStationRepo(db),
		Terminal: postgres.NewTerminalRepo(db),
		Transaction: postgres.NewTransactionRepo(db),
	}
}
