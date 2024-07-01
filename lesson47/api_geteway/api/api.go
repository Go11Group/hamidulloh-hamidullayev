package api

import (
	handler "dodi/api/hendler"
	"dodi/genproto/protos"

	"github.com/gin-gonic/gin"
	// "google.golang.org/grpc"
)

func NewRouter(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	weather := protos.NewWeatherServiceClient(conn)
	transport := protos.NewTransportServiceClient(conn)

	handler := handler.NewHandler(weather, transport)

	router.GET("/get", handler.Get)

	return router
}