package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Database connection constants
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "imthon"
	password = "dodi"
)

// ConnectDB establishes a connection to the PostgreSQL database and returns the connection object.
func ConnectDB() (*sql.DB, error) {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
