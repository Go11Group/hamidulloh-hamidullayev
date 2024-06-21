package postgres

import (
	"database/sql"
	"time"

	"github.com/Go11Group/language/model"

	_ "github.com/lib/pq"
)

// UserRepo represents the repository for managing user data in the database
type UserRepo struct {
	db *sql.DB
}

// NewUserRepo creates a new UserRepo instance
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Get retrieves a user by their ID.
func (u *UserRepo) Get(id string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow("SELECT user_id, name, email, birthday, password, created_at, update_at FROM users WHERE user_id = $1 AND deleted_at = 0", id).Scan(
		&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

// Create adds a new user to the database.
func (u *UserRepo) Create(user model.User) error {
	_, err := u.db.Exec("INSERT INTO users (name, email, birthday, password, deleted_at, update_at) VALUES ($1, $2, $3, $4, $5, $6)", user.Name, user.Email, user.Birthday, user.Password, 0, time.Now())
	return err
}

// Update modifies an existing user's information in the database.
func (u *UserRepo) Update(id string, user model.User) error {
	_, err := u.db.Exec("UPDATE users SET name = $1, email = $2, birthday = $3, password = $4, update_at = $6 WHERE user_id = $5 and deleted_at = 0",
		user.Name, user.Email, time.Now(), user.Password, id, time.Now())
	return err
}

// Delete marks a user as deleted in the database
func (u *UserRepo) Delete(id string) error {
	_, err := u.db.Exec("update users set deleted_at = date_part('epoch', current_timestamp)::INT where user_id = $1 and deleted_at = 0", id)
	return err
}
