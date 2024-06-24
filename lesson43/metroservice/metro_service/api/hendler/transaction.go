package handler

import (
	"fmt"
	"net/http"

	"github.com/Go11Group/lesson43/metro_service/models"
	"github.com/gin-gonic/gin"
)


func (h *handler) CreateTransaction(ctx *gin.Context) {
    transaction := models.Transaction{}

    err := ctx.ShouldBindJSON(&transaction)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    err = h.Transaction.Create(&transaction)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) UpdateTransaction(ctx *gin.Context) {
    id := ctx.Param("id")
    transaction := models.Transaction{}

    err := ctx.ShouldBindJSON(&transaction)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    err = h.Transaction.Update(id, &transaction)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, "Updated")
}

func (h *handler) DeleteTransaction(ctx *gin.Context) {
    id := ctx.Param("id")

    err := h.Transaction.Delete(id)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, "Deleted")
}

func (h *handler) GetTransactionById(ctx *gin.Context) {
    id := ctx.Param("id")

    transaction, err := h.Transaction.GetById(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, err.Error())
        fmt.Println("error:", err.Error())
        return
    }

    ctx.JSON(http.StatusOK, transaction)
}