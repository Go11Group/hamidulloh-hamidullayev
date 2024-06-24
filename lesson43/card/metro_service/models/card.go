package models

type Card struct {
	ID        string `json:"id"`
	Number    string `json:"number"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt int    `json:"deleted_at"`
}
