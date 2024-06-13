package main

import (
	handler "dodi/hendler"
	"dodi/postgres"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	user := postgres.NewUserRepo(db)
	course := postgres.NewCourseRepo(db)
	lesson := postgres.NewLessonRepo(db)
	enrollment := postgres.NewEnrollmentRepo(db)

	h := handler.Handler{
		User:       user,
		Course:     course,
		Lesson:     lesson,
		Enrollment: enrollment,
	}

	router := gin.Default()

	//user
	router.GET("/user/:id", h.UserGet)
	router.POST("/user/create", h.UserCreate)
	router.PUT("/user/update/:id", h.UserUpdate)
	router.DELETE("/user/delete/:id", h.UserDelete)
	router.GET("/user/getall", h.UserGetAll)

	//course
	router.GET("/course/:id", h.CourseGet)
	router.POST("/course/create", h.CourseCreate)
	router.PUT("/course/update/:id", h.CourseUpdate)
	router.DELETE("/course/delete/:id", h.CourseDelete)

	//lesson
	router.GET("/lesson/:id", h.LessonGet)
	router.POST("/lesson/create", h.LessonCreate)
	router.PUT("/lesson/update/:id", h.LessonUpdate)
	router.DELETE("/lesson/delete/:id", h.LessonDelete)

	//enrollment
	router.GET("/enrollment/:id", h.EnrollmentGet)
	router.POST("/enrollment/create", h.EnrollmentCreate)
	router.PUT("/enrollment/update/:id", h.EnrollmentUpdate)
	router.DELETE("/enrollment/delete/:id", h.EnrollmentDelete)

	log.Println("Server started at :8080")
	log.Fatal(router.Run(":8080"))
}
