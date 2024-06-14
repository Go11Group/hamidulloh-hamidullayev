package hendler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Go11Group/language/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) UserGet(c *gin.Context) {
	id := c.Param("id")
	user, err := h.User.Get(id)
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

func (h *Handler) UserSearch(c *gin.Context) {
	name := c.Query("name")
	email := c.Query("email")
	users, err := h.User.SearchUsers(name, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"results": users})
}

func (h *Handler) UserGetAll(c *gin.Context) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset = ""
	)

	query := `select * from users where deleted_at = 0`
	if len(c.Query("name")) > 0 {
		params["name"] = c.Query("name")
		query += " and name = :name"
	}

	if len(c.Query("email")) > 0 {
		params["email"] = c.Query("email")
		query += " and email = :email "
	}

	if len(c.Query("birthday")) > 0 {
		params["birthday"] = c.Query("birthday")
		query += " and birthday = :birthday "
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

	query, arr = UserReplaceQueryParams(query, params)

	users, err := h.User.GetAll(query, arr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ERROR IN GET"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func UserReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
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
