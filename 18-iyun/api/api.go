package routes

import (
    "github.com/gin-gonic/gin"
    "dodi/api/handler"
    "dodi/middlewaare"
)

func SetupRoutes(r *gin.Engine) {
    admin := r.Group("/admin")
    {
        admin.POST("/resource", middleware.AdminAuth(), handlers.CreateResource)
        admin.PUT("/resource/:id", middleware.AdminAuth(), handlers.UpdateResource)
        admin.DELETE("/resource/:id", middleware.AdminAuth(), handlers.DeleteResource)
    }

    r.GET("/resource/:id", handlers.GetResource)
    r.GET("/resources", handlers.ListResources)
}
