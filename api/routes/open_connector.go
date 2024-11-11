package routes

import (
	connector "outlook-connector/api/controllers"
	"outlook-connector/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RoutesOpen(rg *gin.RouterGroup, env config.Config) {
	authController := connector.NewAuthController(env)

	rg.POST("/token", authController.GetToken)
	rg.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
