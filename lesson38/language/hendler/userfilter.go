package hendler

import (
	"dodi/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserGetAll(c *gin.Context) {
	var filter model.UserFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	users, err := h.UserFilter.Users_GetAll(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.JSON(http.StatusOK, users)
}