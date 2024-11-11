package routes

import (
	connector "outlook-connector/api/controllers"
	"outlook-connector/config"

	"github.com/gin-gonic/gin"
)

func RoutesConnector(rg *gin.RouterGroup, env config.Config) {
	connectorController := connector.NewConnectorController(env)
	authController := connector.NewAuthController(env)

	groupLogin := rg.Group("/auth")
	groupLogin.POST("/create", authController.GetToken)

	groupConnector := rg.Group("/connector")
	groupConnector.POST("/read", connectorController.GetEmailFilter)
}
