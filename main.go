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
	router.Static("/build", "./frontend/static/build")
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/static/index.html")
	})

	err = router.Run(":8080")
	errorHandler(err, 2)
}
