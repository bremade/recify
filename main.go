package main

import (
    "github.com/gin-gonic/gin"
    
    "github.com/bremade/recify/api"
)

func main() {
    router := gin.Default()

    // Api
    apiRouter := router.Group("/api/v1")
    apiRouter.GET("/status", api.Status)
    
    // Static files
    router.Static("/build", "./frontend/public/build")
    router.StaticFile("/", "./frontend/public/index.html")

    router.Run(":8000")
}
