package routes

import (
	connector "outlook-connector/api/controllers"
	"outlook-connector/config"

	"github.com/gin-gonic/gin"
)

func RoutesConnector(rg *gin.RouterGroup, env config.Config) {
	connectorController := connector.NewConnectorController(env)

	rg.POST("/read", connectorController.GetEmailFilter)
	rg.POST("/read/full", connectorController.GetEmailFilterFull)
	rg.POST("/mark", connectorController.MarkEmailID)
	rg.POST("/move", connectorController.MoveTo)
	rg.POST("/folders", connectorController.GetFolders)

}
