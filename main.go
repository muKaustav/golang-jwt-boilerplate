package main

import (
	"golang/initializers"
	"golang/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	auth := r.Group("/auth")
	{
		routes.SetupAuthRoutes(auth)
	}

	author := r.Group("/author")
	{
		routes.SetupAuthorRoutes(author)
	}

	r.Run(":" + "8000")
}
