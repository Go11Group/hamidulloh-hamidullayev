package hendler

import (
	"net/http"
	"dodi/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) EnrollmentGet(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}

	enrollment, err := h.Enrollment.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting enrollment"})
		return
	}

	c.JSON(http.StatusOK, enrollment)
}

func (h *Handler) EnrollmentCreate(c *gin.Context) {
	var enrollment model.Enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.Enrollment.Create(enrollment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating enrollment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Enrollment created successfully"})
}

func (h *Handler) EnrollmentUpdate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}

	var updatedEnrollment model.Enrollment
	if err := c.ShouldBindJSON(&updatedEnrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.Enrollment.Update(id, updatedEnrollment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating enrollment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment updated successfully"})
}

func (h *Handler) EnrollmentDelete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}

	if err := h.Enrollment.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting enrollment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment deleted successfully"})
}
