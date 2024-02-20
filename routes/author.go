package routes

import (
	"golang/controllers"
	"golang/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupAuthorRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.AuthorRead)
	r.GET("/me", middlewares.IsAuth, controllers.AuthorMyProfile)
	r.GET("/:id", middlewares.IsAuth, controllers.AuthorReadOne)
	r.PUT("/:id", controllers.AuthorUpdate)
	r.DELETE("/:id", controllers.AuthorDelete)
}
