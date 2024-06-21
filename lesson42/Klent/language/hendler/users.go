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

// UserGet handles the HTTP GET request to retrieve a user by ID.
func (h *Handler) UserGet(c *gin.Context) {
	id := c.Param("id")
	resp, err := http.Get("http://localhost:8080/user/" + id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while decoding JSON"})
		return
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	users := model.User{}

	json.Unmarshal(body, &users)

	c.JSON(http.StatusOK, users)
}

// UserCreate handles the HTTP POST request to create a new user.
func (h *Handler) UserCreate(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := "http://localhost:8080/user/create"

	res, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// UserUpdate handles the HTTP PUT request to update an existing user
func (h *Handler) UserUpdate(c *gin.Context) {
	id := c.Param("id")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	url := "http://localhost:8080/user/update/" + id

	client := http.Client{}
	fmt.Println(body)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// UserDelete handles the HTTP DELETE request to delete a user by ID.
func (h *Handler) UserDelete(c *gin.Context) {
	id := c.Param("id")

	url := "http://localhost:8080/user/delete/" + id

	client := http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}
	_, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
