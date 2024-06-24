package postgres

import (
	"database/sql"

	"github.com/Go11Group/lesson43/metro_service/models"
)

type TransactionRepo struct {
	Db *sql.DB
}

func NewTransactionRepo(db *sql.DB) *TransactionRepo {
	return &TransactionRepo{Db: db}
}

func (tr *TransactionRepo) Create(transaction *models.Transaction) error {
	_, err := tr.Db.Exec("INSERT INTO transaction (card_id, amount, terminal_id, transaction_type) VALUES ($1, $2, $3, $4)",
		transaction.CardID, transaction.Amount, transaction.TerminalID, transaction.TransactionType)
	return err
}

func (tr *TransactionRepo) Update(id string, transaction *models.Transaction) error {
	_, err := tr.Db.Exec("UPDATE transaction SET card_id = $1, amount = $2, terminal_id = $3, transaction_type = $4 WHERE id = $5",
		transaction.CardID, transaction.Amount, transaction.TerminalID, transaction.TransactionType, id)
	return err
}

func (tr *TransactionRepo) Delete(id string) error {
	_, err := tr.Db.Exec("UPDATE transaction SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE id = $1 AND deleted_at = 0", id)
	return err
}

func (tr *TransactionRepo) GetById(id string) (*models.Transaction, error) {
	var transaction = models.Transaction{Id: id}

	err := tr.Db.QueryRow("SELECT card_id, amount, terminal_id, transaction_type FROM transaction WHERE id = $1", id).Scan(&transaction.CardID, &transaction.Amount, &transaction.TerminalID, &transaction.TransactionType)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
