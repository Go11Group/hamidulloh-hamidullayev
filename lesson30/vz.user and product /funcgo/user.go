package funcgo

import (
	"database/sql"
	"Modul/modul"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(DB *sql.DB) *UserRepo {
	return &UserRepo{DB: DB}
}

func (c *UserRepo) UserCreate(user modul.User) error {
	_, err := c.DB.Exec("INSERT INTO users (userName, email, password) VALUES ($1, $2, $3)",
		user.UserName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserRepo) UserDelete(id uint) error {
	_, err := c.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (c *UserRepo) UserUpdate(id uint, user modul.User) error {
	_, err := c.DB.Exec("UPDATE users SET userName = $1, email = $2, password = $3 WHERE id = $4",
		user.UserName, user.Email, user.Password, id)
	if err != nil {
		return err
	}
	return nil
}
func (c *UserRepo) UserID(id uint) (*modul.User, error) {
	user := &modul.User{}
	err := c.DB.QueryRow("SELECT id, userName, email, password FROM users WHERE id = $1", id).
		Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *UserRepo) UserSelect() ([]modul.User, error) {
	rows, err := c.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []modul.User{}
	for rows.Next() {
		user := modul.User{}
		err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
