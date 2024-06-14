package hendler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Go11Group/language/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CourseGet(c *gin.Context) {
	id := c.Param("id")
	course, err := h.Course.Get(id)
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

type UserCourse struct {
	UserId  string         `json:"user_id"`
	Courses []model.Course `json:"courses"`
}

func (h *Handler) GetCoursesByUser(c *gin.Context) {
	userID := c.Param("id")

	courses, err := h.Course.CoursesByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	UserCourse := UserCourse{
		UserId:  userID,
		Courses: courses,
	}
	c.JSON(http.StatusOK, UserCourse)
}

func (h *Handler) CourseGetAll(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from course where deleted_at = 0`
	if len(c.Query("title")) > 0 {
		params["title"] = c.Query("title")
		query += " and title = :title"
	}

	if len(c.Query("description")) > 0 {
		params["description"] = c.Query("description")
		query += " and description = :description "
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

	query, arr = CourseReplaceQueryParams(query, params)

	courses, err := h.Course.GetAll(query, arr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERROR IN GET"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func CourseReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
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
