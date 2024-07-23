package main

import (
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "context"
    "dodi/api"
)

var ctx = context.Background()

func main() {
    r := gin.Default()

    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", 
    })

    r.Use(func(c *gin.Context) {
        c.Set("redis", rdb)
        c.Next()
    })

    routes.SetupRoutes(r)

    r.Run(":8080")
}
