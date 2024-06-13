package hendler

import (
	"dodi/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) LessonGet(c *gin.Context) {
	idStr := c.Param("id")
	lesson, err := h.Lesson.Get(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting enrollment"})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (h *Handler) LessonCreate(c *gin.Context) {
	var lesson model.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}
	if err := h.Lesson.Create(lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating enrollment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Enrollment created successfully"})
}

func (h *Handler) LessonUpdate(c *gin.Context) {
	idStr := c.Param("id")
	var lesson model.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.Lesson.Update(idStr, lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating enrollment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment updated successfully"})
}

func (h *Handler) LessonDelete(c *gin.Context) {
	idStr := c.Param("id")
	if err := h.Lesson.Delete(idStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting enrollment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment deleted successfully"})
}