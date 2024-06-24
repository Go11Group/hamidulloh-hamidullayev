package postgres

import (
	"database/sql"

	"github.com/Go11Group/lesson43/metro_service/models"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) Create(user *models.User) error {
	_, err := u.Db.Exec("INSERT INTO users (name, phone, age) VALUES ($1, $2, $3)",
		user.Name, user.Phone, user.Age)
	return err
}

func (u *UserRepo) Update(id string, user *models.User) error {
	_, err := u.Db.Exec("UPDATE users SET name = $1, phone = $2, age = $3 WHERE id = $4 AND deleted_at = 0",
		user.Name, user.Phone, user.Age, id)
	return err
}

func (t *UserRepo) Delete(id string) error {
	_, err := t.Db.Exec("UPDATE users SET deleted_at = date_part('epoch', current_timestamp)::INT WHERE id = $1 AND deleted_at = 0", id)
	return err
}

func (u *UserRepo) GetById(id string) (*models.User, error) {
	var user = models.User{}

	err := u.Db.QueryRow("SELECT name, phone, age, created_at, updated_at FROM users WHERE id = $1 AND deleted_at = 0", id).
		Scan(&user.Name, &user.Phone, &user.Age, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	user.Id = id
	return &user, nil
}
