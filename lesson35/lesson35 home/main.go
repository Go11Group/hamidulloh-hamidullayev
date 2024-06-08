package main

import (
	"log"
	// "net/http"

	"github.com/Go11Group/at_lesson/lesson34/handler"
	"github.com/Go11Group/at_lesson/lesson34/postgres"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Initialize the problem repository
	problemRepo := postgres.Problem_NewBookRepo(db)

	// Initialize the HTTP server handler with the problem repository
	server := handler.NewHandler(handler.Handler{
		Problem: problemRepo,
	})

	// Start the HTTP server
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}