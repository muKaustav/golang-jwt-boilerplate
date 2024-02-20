package routes

import (
	"golang/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(r *gin.RouterGroup) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
}
