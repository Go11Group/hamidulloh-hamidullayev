package postgres

import (
	"database/sql"
	"dodi/model"
	_ "github.com/lib/pq"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) User_Get(id int) (model.User, error) {
	var user model.User
	err := u.db.QueryRow("SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE user_id = $1", id).Scan(
		&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	return user, err
}

func (u *UserRepo) User_Create(user model.User) error {
	_, err := u.db.Exec("INSERT INTO users (name, email, birthday, password) VALUES ($1, $2, $3, $4)", &user.Name, &user.Email, &user.Birthday, user.Password)
	return err
}

func (u *UserRepo) User_Update(id int, user model.User) error {
	_, err := u.db.Exec("UPDATE users SET name = $1, email = $2, birthday = $3, password = $4 WHERE user_id = $5",
		user.Name, user.Email, user.Birthday, user.Password, id)
	return err
}

func (u *UserRepo) User_Delete(id int) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (us *UserRepo) GetAllUser() ([]model.User, error) {
	rows, err := us.db.Query("SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		user := model.User{}
		var birthday, deletedAt string
		err = rows.Scan(&user.UserID, &user.Name, &user.Email, &birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &deletedAt)
		if err != nil {
			return nil, err
		}
	}
	return users, nil
}