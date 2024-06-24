package postgres

import (
	"database/sql"

	"github.com/Go11Group/lesson43/metro_service/models"
)

type TerminalRepo struct {
	Db *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{Db: db}
}

func (t *TerminalRepo) Create(terminal *models.Terminal) error {
	_, err := t.Db.Exec("INSERT INTO terminal (station_id) VALUES ($1)", terminal.StationID)
	return err
}

func (t *TerminalRepo) Update(id string, terminal *models.Terminal) error {
	_, err := t.Db.Exec("UPDATE terminal SET station_id = $1 WHERE id = $2",
		terminal.StationID, id)
	return err
}

func (t *TerminalRepo) Delete(id string) error {
	_, err := t.Db.Exec("UPDATE terminal SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE id = $1 AND deleted_at = 0", id)
	return err
}

func (t *TerminalRepo) GetById(id string) (*models.Terminal, error) {
	var terminal = models.Terminal{Id: id}

	err := t.Db.QueryRow("SELECT station_id FROM terminal WHERE id = $1", id).Scan(&terminal.StationID)
	if err != nil {
		return nil, err
	}

	return &terminal, nil
}
