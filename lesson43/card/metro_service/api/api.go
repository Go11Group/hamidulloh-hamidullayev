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

	r.POST("/cards", h.CreateCard)
	r.PUT("/cards/:id", h.UpdateCard)
	r.DELETE("/cards/:id", h.DeleteCard)
	r.GET("/cards/:id", h.GetCardById)

    
	return &http.Server{Handler: r, Addr: ":8082"}
}
