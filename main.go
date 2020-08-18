package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    
    "github.com/bremade/recify/api"
    "github.com/bremade/recify/auth"
    "github.com/bremade/recify/persistence"
)

func errorHandler(err error, exitCode int) {
	if err != nil {
		os.Exit(exitCode)
	}
}

func main() {
	router := gin.Default()
    router.Use(sessions.Sessions("recify_session", auth.GetSessionStore()))

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}

	// Database
	db := persistence.DB{}

	dbURI := os.Getenv("DATABASE_URI")
	if len(dbURI) == 0 {
		panic("not database uri to connect to")
	}

	err := db.Open(dbURI)
	errorHandler(err, 1)

	db.Setup()

	// Api
    recifyApi := api.New(&db)
	apiRouter := router.Group("/api/v1")
	apiRouter.GET("/status", recifyApi.Status)

    // Authentication
    apiRouter.POST("/auth/login", recifyApi.Login)
    apiRouter.GET("/auth/status", recifyApi.AuthStatus)
    apiRouter.POST("/auth/register", recifyApi.Register)
    apiRouter.POST("/auth/logout", recifyApi.Logout)

	// Static files
	router.Static("/build", "./frontend/public/build")
	router.StaticFile("/", "./frontend/public/index.html")

	err = router.Run(":" + port)
	errorHandler(err, 2)
}
