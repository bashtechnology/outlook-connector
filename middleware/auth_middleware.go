package middleware

import (
	"fmt"
	"net/http"
	"outlook-connector/api/data/response"
	"outlook-connector/api/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func HeaderAuthBearer(SecretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		authHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authHeader)
		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}
		if token == "" {
			webResponse := response.HttpResponse{
				Code:    http.StatusUnauthorized,
				Status:  "Bad Request",
				Message: fmt.Sprintf("Token Invalido."),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
			return
		}
		id, err := utils.ValidateToken(token, SecretKey)
		if err != nil {
			webResponse := response.HttpResponse{
				Code:    http.StatusUnauthorized,
				Status:  "Bad Request",
				Message: fmt.Sprintf("Falha ao validar o Token."),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
			return
		}
		c.Set("currentUser", id)
		c.Next()
	}
}
func HeaderAuthToken(SecretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		authHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authHeader)
		if len(fields) != 0 && fields[0] == "Token" {
			token = fields[1]
		}
		if token == "" {
			webResponse := response.HttpResponse{
				Code:    http.StatusUnauthorized,
				Status:  "Bad Request",
				Message: fmt.Sprintf("Token inválido."),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
			return
		}
		if token != SecretKey {
			webResponse := response.HttpResponse{
				Code:    http.StatusUnauthorized,
				Status:  "Bad Request",
				Message: fmt.Sprintf("Token Não Autorizado."),
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, webResponse)
			return
		}
		c.Next()
	}
}
