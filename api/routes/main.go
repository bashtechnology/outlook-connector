package routes

import (
	"net/http"
	"outlook-connector/config"
	"outlook-connector/middleware"
	"strings"

	_ "outlook-connector/docs"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func Init(ctx config.Config) *gin.Engine {
	middleware.ConfigurarCORSMiddleware(router, ctx.AllowedOrigins)
	getRoutes(ctx)
	return router
}

func getRoutes(env config.Config) {
	router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/api/open/docs/") {
			c.Redirect(http.StatusTemporaryRedirect, "/api/open/docs/swagger/index.html")
			return
		}
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Recurso não encontrado."})
	})
	api := router.Group("/api")
	v1 := api.Group("/v1")
	open := v1.Group("/open")
	connector := v1.Group("/connector")
	connector.Use(middleware.HeaderAuthBearer(env.TokenSecret))
	RoutesConnector(connector, env)
	RoutesOpen(open, env)
}
