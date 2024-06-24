package handler

import (
	"net/http"

	"github.com/Go11Group/lesson43/metro_service/models"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCard(ctx *gin.Context) {
	card := models.Card{}

	if err := ctx.ShouldBindJSON(&card); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Card.Create(&card); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, card)
}

func (h *handler) UpdateCard(ctx *gin.Context) {
	id := ctx.Param("id")
	card := models.Card{}

	if err := ctx.ShouldBindJSON(&card); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Card.Update(id, &card); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, card)
}

func (h *handler) DeleteCard(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := h.Card.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Card deleted"})
}

func (h *handler) GetCardById(ctx *gin.Context) {
	id := ctx.Param("id")

	card, err := h.Card.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, card)
}