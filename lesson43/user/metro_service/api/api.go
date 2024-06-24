package api

import (
	"database/sql"
	"net/http"

	handler "github.com/Go11Group/lesson43/metro_service/api/hendler"
	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *http.Server {
	r := gin.Default()

	h := handler.NewHandler(db)

	r.POST("/users", h.CreateUser)
    r.PUT("/users/:id", h.UpdateUser)
    r.DELETE("/users/:id", h.DeleteUser)
    r.GET("/users/:id", h.GetUserById)

    
	return &http.Server{Handler: r, Addr: ":8081"}
}
