package handler

import (
	"fmt"
	"net/http"

	"github.com/Go11Group/lesson43/metro_service/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateStation(ctx *gin.Context) {
	stn := models.Station{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Station.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) UpdateStation(ctx *gin.Context) {
    id := ctx.Param("id")
    stn := models.Station{}

    err := ctx.ShouldBindJSON(&stn)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    err = h.Station.Update(id, &stn)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, "Updated")
}

func (h *handler) DeleteStation(ctx *gin.Context) {
    id := ctx.Param("id")

    err := h.Station.Delete(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, "Deleted")
}

func (h *handler) GetStationById(ctx *gin.Context) {
    id := ctx.Param("id")

    station, err := h.Station.GetById(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, station)
}