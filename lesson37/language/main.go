package main

import (
	handler "dodi/hendler"
	"dodi/postgres"
	"log"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	user := postgres.NewUserRepo(db)
	course := postgres.NewCourseRepo(db)

	h := handler.Handler{
		User:          user,
		Course:       course,
	}

	router := gin.Default()

	// User routes
	router.GET("/user/:id", h.Userr)
	router.POST("/user/create", h.UserCreate)
	router.PUT("/user/update/:id", h.UserUpdate)
	router.DELETE("/user/delete/:id", h.UserDelete)

	router.GET("/course/:id",h.Course_Get)
	router.POST("/course/create", h.CourseCreate)
	router.PUT("/course/update/:id", h.CourseUpdate)
	router.DELETE("/course/delete/:id", h.CourseDelete)


	log.Println("Server started at :8080")
	log.Fatal(router.Run(":8080"))
}