package postgres

import (
	"database/sql"

	"github.com/Go11Group/lesson43/metro_service/models"
)

type CardRepo struct {
	Db *sql.DB
}

func NewCardRepo(db *sql.DB) *CardRepo {
	return &CardRepo{Db: db}
}

func (c *CardRepo) Create(card *models.Card) error {
	_, err := c.Db.Exec("INSERT INTO cards (number, user_id) VALUES ($1, $2)",
		card.Number, card.UserID)
	return err
}

func (c *CardRepo) Update(id string, card *models.Card) error {
	_, err := c.Db.Exec("UPDATE cards SET number = $1, user_id = $2 WHERE id = $3 AND deleted_at = 0",
		card.Number, card.UserID, id)
	return err
}

func (t *CardRepo) Delete(id string) error {
	_, err := t.Db.Exec("UPDATE cards SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE id = $1 AND deleted_at = 0", id)
	return err
}


func (c *CardRepo) GetById(id string) (*models.Card, error) {
	card := &models.Card{}

	err := c.Db.QueryRow("SELECT number, user_id, created_at, updated_at FROM cards WHERE id = $1 AND deleted_at = 0", id).
		Scan(&card.Number, &card.UserID,  &card.CreatedAt, &card.UpdatedAt)
	if err != nil {
		return nil, err
	}

	card.ID = id
	return card, nil
}
