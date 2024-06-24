package postgres

import (
	"database/sql"

	"github.com/Go11Group/lesson43/metro_service/models"
)

type StationRepo struct {
	Db *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{Db: db}
}

func (s *StationRepo) Create(station *models.Station) error {

	_, err := s.Db.Exec("insert into station(name) values ($1)", station.Name)

	return err
}

func (s *StationRepo) Update(id string, station *models.Station) error {
	_, err := s.Db.Exec("UPDATE station SET name = $1 WHERE id = $2",
		station.Name, id)
	return err
}

func (s *StationRepo) Delete(id string) error {
	_, err := s.Db.Exec("UPDATE station SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE id = $1 AND deleted_at = 0", id)
	return err
}

func (s *StationRepo) GetById(id string) (*models.Station, error) {
	var station = models.Station{Id: id}

	err := s.Db.QueryRow("select name from station where id = $1", id).
		Scan(&station.Name)
	if err != nil {
		return nil, err
	}

	return &station, nil
}
