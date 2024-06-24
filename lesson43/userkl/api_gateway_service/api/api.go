package api

import (
	"net/http"

	handler "github.com/Go11Group/lesson43/api_gateway_service/api/hendler"
	"github.com/gin-gonic/gin"
)

func Routes() *http.Server {
	mux := gin.Default()

	h := handler.NewHandler()

	mux.GET("/stations/:id", h.Client)
	mux.POST("/stations", h.Client)
	mux.PUT("/stations/:id", h.Client)
	mux.DELETE("/stations/:id", h.Client)

	mux.POST("/terminals", h.Client)
	mux.PUT("/terminals/:id", h.Client)
	mux.DELETE("/terminals/:id", h.Client)
	mux.GET("/terminals/:id", h.Client)

	mux.POST("/transactions", h.Client)
	mux.PUT("/transactions/:id", h.Client)
	mux.DELETE("/transactions/:id", h.Client)
	mux.GET("/transactions/:id", h.Client)

	mux.POST("/users", h.UserClient)
	mux.PUT("/users/:id", h.UserClient)
	mux.DELETE("/users/:id", h.UserClient)
	mux.GET("/users/:id", h.UserClient)

	mux.POST("/cards", h.CardClient)
	mux.PUT("/cards/:id", h.CardClient)
	mux.DELETE("/cards/:id", h.CardClient)
	mux.GET("/cards/:id", h.CardClient)

	return &http.Server{Handler: mux, Addr: ":8083"}
}
