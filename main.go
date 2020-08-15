package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/bremade/recify/api"
	"github.com/bremade/recify/persistence"
)

func errorHandler(err error, exitCode int) {
	if err != nil {
		os.Exit(exitCode)
	}
}

func main() {
	router := gin.Default()

	// Database
	db := persistence.DB{}

	dbURI := os.Getenv("DATABASE_URI")
	if len(dbURI) == 0 {
		panic("not database uri to connect to")
	}

	err := db.Open(dbURI)
	errorHandler(err, 1)

	err = db.Setup()
	errorHandler(err, 1)

	_, err = db.QueryTest(1)
	errorHandler(err, 1)

	// Api
	apiRouter := router.Group("/api/v1")
	apiRouter.GET("/status", api.Status)

	// Static files
	router.Static("/build", "./frontend/public/build")
	router.StaticFile("/", "./frontend/public/index.html")

	err = router.Run(":8080")
	errorHandler(err, 2)
}
