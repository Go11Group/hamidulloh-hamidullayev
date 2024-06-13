package postgres

import (
	"database/sql"
	"dodi/model"
	"fmt"
	"strings"
	_ "github.com/lib/pq"
)

type UserFilterRepo struct {
	db *sql.DB
}

func NewUserFilterRepo(db *sql.DB) *UserFilterRepo {
	return &UserFilterRepo{db: db}
}

func (u UserFilterRepo) Users_GetAll(filter model.UserFilter) ([]model.User, error) {
	var users []model.User
	var params []string
	var args []interface{}

	query := `
		SELECT user_id, name, email, birthday, created_at, updated_at
		FROM users WHERE deleted_at IS NULL
	`

	if filter.Name != nil {
		args = append(args, *filter.Name)
		params = append(params, fmt.Sprintf("name = $%d", len(args)))
	}

	if filter.Email != nil {
		params = append(params, fmt.Sprintf("email = $%d", len(args)+1))
		args = append(args, *filter.Email)
	}

	if filter.Birthday != nil {
		params = append(params, fmt.Sprintf("birthday = $%d", len(args)+1))
		args = append(args, filter.Birthday)
	}

	if filter.CreatedAt != nil {
		params = append(params, fmt.Sprintf("created_at = $%d", len(args)+1))
		args = append(args, *filter.CreatedAt)
	}

	if filter.UpdatedAt != nil {
		params = append(params, fmt.Sprintf("updated_at = $%d", len(args)+1))
		args = append(args, *filter.UpdatedAt)
	}

	if len(params) > 0 {
		query += strings.Join(params, " AND ")
	}

	if filter.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *filter.Limit)
	}

	if filter.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *filter.Offset)
	}

	rows, err := u.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed while querying: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Birthday, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed while scanning data to slice: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return users, nil
}