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

    r.POST("/stations", h.CreateStation)
    r.PUT("/stations/:id", h.UpdateStation)
	r.DELETE("/stations/:id", h.DeleteStation)
	r.GET("/stations/:id", h.GetStationById)

	r.POST("/terminals", h.CreateTerminal)
    r.PUT("/terminals/:id", h.UpdateTerminal)
    r.DELETE("/terminals/:id", h.DeleteTerminal)
    r.GET("/terminals/:id", h.GetTerminalById)

	r.POST("/transactions", h.CreateTransaction)
    r.PUT("/transactions/:id", h.UpdateTransaction)
    r.DELETE("/transactions/:id", h.DeleteTransaction)
    r.GET("/transactions/:id", h.GetTransactionById)

	return &http.Server{Handler: r, Addr: ":8080"}
}
