package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func ValidateQueryParams(queryParams map[string]string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Extrair os parâmetros da consulta da solicitação
		params := ctx.Request.URL.Query()

		// Validar os parâmetros de consulta dinamicamente
		validate := validator.New()
		for paramName, validationRule := range queryParams {
			paramValue := params.Get(paramName)
			if err := validate.Var(paramValue, validationRule); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"message": "Parâmetro de consulta inválido: " + paramName})
				ctx.Abort()
				return
			}
		}

		// Os parâmetros de consulta são válidos, continue com o processamento
		ctx.Next()
	}
}
