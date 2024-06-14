package hendler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Go11Group/language/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) EnrollmentGet(c *gin.Context) {
	id := c.Param("id")
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

	c.JSON(http.StatusCreated, gin.H{"message": " Enrollment created successfully"})
}

func (h *Handler) EnrollmentUpdate(c *gin.Context) {
	id := c.Param("id")
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
	id := c.Param("id")
	if err := h.Enrollment.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting enrollment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment deleted successfully"})
}

func (h *Handler) GetEnrolledUsersByCourse(c *gin.Context) {
	id := c.Param("id")
	enrolledUsers, err := h.Enrollment.EnrolledUsersByCourse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrolled users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course_id": id, "enrolled_users": enrolledUsers})
}

func (h *Handler) EnrollmentGetAll(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from lessons where deleted_at = 0`
	if len(c.Query("title")) > 0 {
		params["title"] = c.Query("title")
		query += " and title = :title"
	}

	if len(c.Query("content")) > 0 {
		params["content"] = c.Query("content")
		query += " and content = :content "
	}
	if len(c.Query("course_id")) > 0 {
		params["course_id"] = c.Query("course_id")
		query += " and course_id = :course_id "
	}


	if len(c.Query("limit")) > 0 {
		params["limit"] = c.Query("limit")
		limit = ` LIMIT :limit`
	}

	if len(c.Query("offset")) > 0 {
		params["offset"] = c.Query("offset")
		offset = ` OFFSET :offset`
	}

	query = query + limit + offset

	query, arr = EnrollmentReplaceQueryParams(query, params)

	lessons, err := h.Lesson.GetAll(query, arr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERROR IN GET"})
		return
	}
	c.JSON(http.StatusOK, lessons)
}

func EnrollmentReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	var (
		i    int = 1
		args []interface{}
	)
	for k, v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}
	return namedQuery, args
}
