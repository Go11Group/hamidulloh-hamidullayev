package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func AdminAuth() gin.HandlerFunc {
    return func(c *gin.Context) {
        if c.GetHeader("Admin") != "true" {
            c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
            c.Abort()
            return
        }
        c.Next()
    }
}
