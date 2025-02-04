package main

import (
	"go-gin-workshop/config"
	"go-gin-workshop/controllers"
	"go-gin-workshop/docs"
	"go-gin-workshop/entities"
	"go-gin-workshop/repositories"
	"go-gin-workshop/routes"
	"go-gin-workshop/services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	//swagger
	docs.SwaggerInfo.BasePath = "/api"

	//db
	db := config.DatabaseConnection()

	db.AutoMigrate(&entities.UserEntity{})

	//repositories
	authRepository := repositories.NewAuthRepository(db)

	//service
	authService := services.NewAuthService(authRepository)

	//controller

	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	routes.NewAuthRoute(r, authController)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
