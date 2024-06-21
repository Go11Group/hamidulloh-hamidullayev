package main

import (
	"log"

	handler "github.com/Go11Group/language/hendler"
	"github.com/Go11Group/language/postgres"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to the database
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Initialize repositories
	user := postgres.NewUserRepo(db)
	course := postgres.NewCourseRepo(db)

	// Initialize handler with repositories
	h := handler.Handler{
		User:   user,
		Course: course,
	}

	// Initialize Gin router
	router := gin.Default()

	//User routes
	router.GET("/user/:id", h.UserGet)
	router.POST("/user/create", h.UserCreate)
	router.PUT("/user/update/:id", h.UserUpdate)
	router.DELETE("/user/delete/:id", h.UserDelete)

	//Course routes
	router.GET("/course/:id", h.CourseGet)
	router.POST("/course/create", h.CourseCreate)
	router.PUT("/course/update/:id", h.CourseUpdate)
	router.DELETE("/course/delete/:id", h.CourseDelete)

	// Start the server
	log.Fatal(router.Run(":8081"))
}
