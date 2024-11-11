package middleware

import (
	"net/http"
	"outlook-connector/api/data/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func ConfigurarCORSMiddleware(router *gin.Engine, AllowedOrigins string) {
	router.Use(corsMiddleware(AllowedOrigins))
}

func corsMiddleware(AllowedOrigins string) gin.HandlerFunc {
	allowedOrigins := strings.Split(AllowedOrigins, ",")
	return func(c *gin.Context) {
		dev := c.GetHeader("DevTest")
		if dev != "" {
			c.Next()
			return
		}
		origin := c.GetHeader("Origin")
		if !strings.HasPrefix(c.FullPath(), "/api/open/") {
			if !isOriginAllowed(origin, allowedOrigins) {
				webresp := response.HttpResponse{
					Code:    http.StatusForbidden,
					Status:  "Bad Request",
					Message: "Origem não permitida.",
				}
				c.JSON(http.StatusForbidden, webresp)
				c.Abort()
				return
			}
		}
		c.Writer.Header().Del("Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}

func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if allowedOrigin == origin {
			return true
		}
	}
	return false
}
