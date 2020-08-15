package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    
    "github.com/bremade/recify/api"
    "github.com/bremade/recify/auth"
)

func main() {
    router := gin.Default()
    router.Use(sessions.Sessions("recify_session", auth.GetSessionStore()))

    // Api
    apiRouter := router.Group("/api/v1")
    apiRouter.GET("/status", api.Status)

    // Authentication
    apiRouter.POST("/auth/login", api.Login)
    apiRouter.GET("/auth/status", api.AuthStatus)
    
    // Static files
    router.Static("/build", "./frontend/public/build")
    router.StaticFile("/", "./frontend/public/index.html")

    router.Run(":8000")
}
