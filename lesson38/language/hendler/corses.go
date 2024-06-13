package hendler

import (
	"net/http"
	"strconv"

	"dodi/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CourseGet(c *gin.Context) {
    idStr := c.Param("id")
    course, err := h.Course.Get(idStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting course"})
        return
    }

    c.JSON(http.StatusOK, course)
}

func (h *Handler) CourseCreate(c *gin.Context) {
    var course model.Course
    if err := c.ShouldBindJSON(&course); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
        return
    }

    if err := h.Course.Create(course); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating course"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Course created successfully"})
}

func (h *Handler) CourseUpdate(c *gin.Context) {
    idStr := c.Param("id")
    var updatedCourse model.Course
    if err := c.ShouldBindJSON(&updatedCourse); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
        return
    }

    if err := h.Course.Update(idStr, updatedCourse); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating course"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

func (h *Handler) CourseDelete(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    if err := h.Course.Delete(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting course"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}