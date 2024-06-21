package hendler

import (
	"net/http"

	"github.com/Go11Group/language/model"
	"github.com/gin-gonic/gin"
)

// CourseGet handles the GET request to retrieve a course by ID.
func (h *Handler) CourseGet(c *gin.Context) {
	id := c.Param("id")
	course, err := h.Course.Get(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting course"})
		return
	}

	c.JSON(http.StatusOK, course)
}

// CourseCreate handles the POST request to create a new course.
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

// CourseUpdate handles the PUT request to update an existing course.
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

// CourseDelete handles the DELETE request to delete a course by ID.
func (h *Handler) CourseDelete(c *gin.Context) {
	id := c.Param("id")

	if err := h.Course.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
