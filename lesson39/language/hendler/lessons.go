package hendler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Go11Group/language/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LessonGet(c *gin.Context) {
	idStr := c.Param("id")
	lesson, err := h.Lesson.Get(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting lesson"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating lesson"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "lesson created successfully"})
}

func (h *Handler) LessonUpdate(c *gin.Context) {
	idStr := c.Param("id")
	var lesson model.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	if err := h.Lesson.Update(idStr, lesson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while updating lesson"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "lesson updated successfully"})
}

func (h *Handler) LessonDelete(c *gin.Context) {
	idStr := c.Param("id")
	if err := h.Lesson.Delete(idStr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while deleting lesson"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "lesson deleted successfully"})
}

func (h *Handler) GetLessonsByCourse(c *gin.Context) {
	id := c.Param("id")
	lessons, err := h.Lesson.LessonsByCourse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting lessons by course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course_id": id, "lessons": lessons})
}

func (h *Handler) LessonGetAll(c *gin.Context) {
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

	query, arr = LessonReplaceQueryParams(query, params)

	lessons, err := h.Lesson.GetAll(query, arr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERROR IN GET"})
		return
	}
	c.JSON(http.StatusOK, lessons)
}

func LessonReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
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
