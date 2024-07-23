package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "net/http"
    "context"
)

var ctx = context.Background()

type Resource struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Value string `json:"value"`
}

func CreateResource(c *gin.Context) {
    var resource Resource
    if err := c.ShouldBindJSON(&resource); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    rdb := c.MustGet("redis").(*redis.Client)
    err := rdb.HSet(ctx, "resource:"+resource.ID, map[string]interface{}{
        "name":  resource.Name,
        "value": resource.Value,
    }).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, resource)
}

func UpdateResource(c *gin.Context) {
    id := c.Param("id")
    var resource Resource
    if err := c.ShouldBindJSON(&resource); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    rdb := c.MustGet("redis").(*redis.Client)
    err := rdb.HSet(ctx, "resource:"+id, map[string]interface{}{
        "name":  resource.Name,
        "value": resource.Value,
    }).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, resource)
}

func DeleteResource(c *gin.Context) {
    id := c.Param("id")
    rdb := c.MustGet("redis").(*redis.Client)
    err := rdb.Del(ctx, "resource:"+id).Err()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Resource deleted"})
}

func GetResource(c *gin.Context) {
    id := c.Param("id")
    rdb := c.MustGet("redis").(*redis.Client)

    resourceMap, err := rdb.HGetAll(ctx, "resource:"+id).Result()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if len(resourceMap) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
        return
    }

    resource := Resource{
        ID:    id,
        Name:  resourceMap["name"],
        Value: resourceMap["value"],
    }

    c.JSON(http.StatusOK, resource)
}

func ListResources(c *gin.Context) {
    rdb := c.MustGet("redis").(*redis.Client)

    keys, err := rdb.Keys(ctx, "resource:*").Result()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    var resources []Resource
    for _, key := range keys {
        resourceMap, err := rdb.HGetAll(ctx, key).Result()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        id := key[len("resource:"):]
        resources = append(resources, Resource{
            ID:    id,
            Name:  resourceMap["name"],
            Value: resourceMap["value"],
        })
    }

    c.JSON(http.StatusOK, resources)
}
