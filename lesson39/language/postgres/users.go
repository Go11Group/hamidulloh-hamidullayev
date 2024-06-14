package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Go11Group/language/model"

	_ "github.com/lib/pq"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow("SELECT user_id, name, email, birthday, password, created_at, update_at FROM users WHERE user_id = $1 AND deleted_at = 0", id).Scan(
		&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (u *UserRepo) Create(user model.User) error {
	_, err := u.db.Exec("INSERT INTO users (name, email, birthday, password, deleted_at, update_at) VALUES ($1, $2, $3, $4, $5, $6)", user.Name, user.Email, user.Birthday, user.Password, 0, time.Now())
	return err
}

func (u *UserRepo) Update(id string, user model.User) error {
	_, err := u.db.Exec("UPDATE users SET name = $1, email = $2, birthday = $3, password = $4, update_at = $6 WHERE user_id = $5 and deleted_at = 0",
		user.Name, user.Email, time.Now(), user.Password, id, time.Now())
	return err
}

func (u *UserRepo) Delete(id string) error {
	_, err := u.db.Exec("update users set deleted_at = date_part('epoch', current_timestamp)::INT where user_id = $1 and deleted_at = 0", id)
	return err
}
func (u *UserRepo) GetAll(a string, arr []interface{}) ([]model.User, error) {
	rows, err := u.db.Query(a, arr...)
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday,
			&user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepo) SearchUsers(name string, email string) ([]model.User, error) {
	query := `SELECT user_id, name, email, birthday, password, created_at, update_at, deleted_at
              FROM users
              WHERE deleted_at = 0`

	args := []interface{}{}
	argIdx := 1

	if name != "" {
		query += fmt.Sprintf(" AND name LIKE $%d", argIdx)
		args = append(args, "%"+name+"%")
		argIdx++
	}
	if email != "" {
		query += fmt.Sprintf(" AND email LIKE $%d", argIdx)
		args = append(args, "%"+email+"%")
		argIdx++
	}

	rows, err := u.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
