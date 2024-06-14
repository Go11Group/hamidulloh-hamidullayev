package main

import (
	"log"

	handler "github.com/Go11Group/language/hendler"
	"github.com/Go11Group/language/postgres"

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
	router.GET("/user/:id", h.UserGet)              //1
	router.GET("/user/getAll",h.UserGetAll)
	router.GET("/users/search", h.UserSearch)       //1
	router.POST("/user/create", h.UserCreate)       //1
	router.PUT("/user/update/:id", h.UserUpdate)    //1
	router.DELETE("/user/delete/:id", h.UserDelete) //1

	//course
	router.GET("/course/:id", h.CourseGet)                //1
	router.GET("/course/getAll",h.CourseGetAll)
	router.GET("/get/courses/by/:id", h.GetCoursesByUser) //1
	router.POST("/course/create", h.CourseCreate)         //1
	router.PUT("/course/update/:id", h.CourseUpdate)      //1
	router.DELETE("/course/delete/:id", h.CourseDelete)   //1

	//lesson
	router.GET("/lesson/:id", h.LessonGet)                 //1
	router.GET("/lesson/getAll",h.LessonGetAll)
	router.GET("/get/lesson/by/:id", h.GetLessonsByCourse) //1
	router.POST("/lesson/create", h.LessonCreate)          //1
	router.PUT("/lesson/update/:id", h.LessonUpdate)       //1
	router.DELETE("/lesson/delete/:id", h.LessonDelete)    //1

	//enrollment
	router.GET("/enrollment/:id", h.EnrollmentGet)                               //1
	router.GET("/enrollment/getAll",h.EnrollmentGetAll)
	router.GET("/courses/course_id/enrollments/:id", h.GetEnrolledUsersByCourse) //1
	router.POST("/enrollment/create", h.EnrollmentCreate)                        //1
	router.PUT("/enrollment/update/:id", h.EnrollmentUpdate)                     //1
	router.DELETE("/enrollment/delete/:id", h.EnrollmentDelete)                  //1

	log.Println("Server started at :8080")
	log.Fatal(router.Run(":8080"))
}
