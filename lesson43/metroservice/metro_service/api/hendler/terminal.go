package handler

import (
	"fmt"
	"net/http"

	"github.com/Go11Group/lesson43/metro_service/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateTerminal(ctx *gin.Context) {
    terminal := models.Terminal{}

    err := ctx.ShouldBindJSON(&terminal)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    err = h.Terminal.Create(&terminal)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) UpdateTerminal(ctx *gin.Context) {
    id := ctx.Param("id")
    terminal := models.Terminal{}

    err := ctx.ShouldBindJSON(&terminal)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    err = h.Terminal.Update(id, &terminal)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, "Updated")
}

func (h *handler) DeleteTerminal(ctx *gin.Context) {
    id := ctx.Param("id")

    err := h.Terminal.Delete(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, "Deleted")
}

func (h *handler) GetTerminalById(ctx *gin.Context) {
    id := ctx.Param("id")

    terminal, err := h.Terminal.GetById(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, terminal)
}