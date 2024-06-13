package hendler

import (
	"dodi/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserGet(c *gin.Context) {
	idStr := c.Param("id")
	user, err := h.User.Get(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UserCreate(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}
	if err := h.User.Create(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *Handler) UserUpdate(c *gin.Context) {
	idStr := c.Param("id")
	var updatedUser model.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.User.Update(idStr, updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (h *Handler) UserDelete(c *gin.Context) {
	idStr := c.Param("id")
	if err := h.User.Delete(idStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
