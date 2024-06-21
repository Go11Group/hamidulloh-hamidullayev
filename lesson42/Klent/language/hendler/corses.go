package hendler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Go11Group/language/model"
	"github.com/gin-gonic/gin"
)

// CourseGet handles the GET request to retrieve a course by ID.
func (h *Handler) CourseGet(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8080/course/" + id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	courses := model.Course{}

	json.Unmarshal(body, &courses)

	c.JSON(http.StatusOK, courses)
}

// CourseCreate handles the POST request to create a new course.
func (h *Handler) CourseCreate(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := "http://localhost:8080/course/create"

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	c.JSON(http.StatusCreated, gin.H{"message": "Course created successfully"})
}

// CourseUpdate handles the PUT request to update an existing course.
func (h *Handler) CourseUpdate(c *gin.Context) {
	id := c.Param("id")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := "http://localhost:8080/course/update/" + id

	client := http.Client{}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

// CourseDelete handles the DELETE request to delete a course by ID.
func (h *Handler) CourseDelete(c *gin.Context) {
	id := c.Param("id")

	url := "http://localhost:8080/course/delete/" + id

	client := http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
