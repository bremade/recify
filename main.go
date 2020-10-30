package main

import (
    "os"
    "net/url"
    "net/http"
    "net/http/httputil"
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

func handleKeycloakProxy(c *gin.Context) {
    remote, err := url.Parse("http://keycloak:8080")
	if err != nil {
		panic(err)
	}

    proxy := httputil.NewSingleHostReverseProxy(remote)
    
    director := func(req *http.Request) {
        req.Header = c.Request.Header
        req.Host = remote.Host
        req.URL.Scheme = remote.Scheme
        req.URL.Host = remote.Host
    }

    proxy.Director = director
    proxy.ServeHTTP(c.Writer, c.Request)
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

    // OAuth2
    providerURL := "http://localhost:8000/auth/realms/recify"

	// Api
    recifyApi := api.New(&db, providerURL)
	apiRouter := router.Group("/api/v1")
	apiRouter.GET("/status", recifyApi.Status)

    // Authentication
    apiRouter.GET("/auth/token", recifyApi.CheckCode)
    apiRouter.GET("/auth/status", recifyApi.AuthStatus)
    
    router.GET("/login", recifyApi.Login)
    router.GET("/logout", recifyApi.Logout)
    
    // Recipe
    apiRouter.GET("/recipe", recifyApi.RetrieveRecipes)
    apiRouter.GET("/recipe/:id", recifyApi.RetrieveRecipe)
    apiRouter.POST("/recipe", recifyApi.CreateRecipe)
    apiRouter.PUT("/recipe", recifyApi.ReplaceRecipe)
    apiRouter.DELETE("/recipe", recifyApi.DeleteRecipe)
    apiRouter.GET("/ingredient", recifyApi.RetrieveIngredients)
    apiRouter.GET("/tag", recifyApi.RetrieveTags)

    // Keycloak
    router.Any("/auth/*path", handleKeycloakProxy)

	// Static files
	router.Static("/build", "./frontend/static/build")
	router.NoRoute(func(c *gin.Context) {
		c.File("./frontend/static/index.html")
	})

	err = router.Run(":" + port)
	errorHandler(err, 2)
}
