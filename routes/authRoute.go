package routes

import (
	"go-gin-workshop/controllers"

	"github.com/gin-gonic/gin"
)

func NewAuthRoute(route *gin.Engine, authController controllers.AuthController) {
	routes := route.Group("/api/auth")
	{
		routes.POST("/register", authController.RegisterController)
	}

}
