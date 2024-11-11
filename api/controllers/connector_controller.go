package connector

import (
	"log/slog"
	"net/http"
	"outlook-connector/api/data/request"
	"outlook-connector/api/data/response"
	"outlook-connector/api/service"
	"outlook-connector/config"

	"github.com/gin-gonic/gin"
)

type ConnectorController struct {
	connectorService *service.ConnectorServiceImpl
}

func NewConnectorController(env config.Config) *ConnectorController {
	connectorService, err_init := service.NewConnectorServiceImpl(env)
	if err_init != nil {
		slog.Debug("⚙️", "Status", "💥Erro ao iniciar Service.", "Error", err_init)
	}
	return &ConnectorController{connectorService: connectorService}
}

// GetEmailFilter godoc
// @Tags Email
// @Summary Leitura de Emails
// @Description Obter os emails para os parametros informados.
// @ID GetEmailFilter
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @In header
// @Name Authorization
// @Type apiKey
// @Success 200 {object} response.HttpResponse{data=response.MonitoramentoResponse} "Dados recebidos!"
// @Failure 400 {object} response.HttpResponse "Requisição Inválida"
// @Router /v1/connector/read [post]
func (c *ConnectorController) GetEmailFilter(ctx *gin.Context) {
	req := request.GetEmailFilterRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		webResponse := response.HttpResponse{
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Parametros invalidos.",
		}
		ctx.JSON(webResponse.Code, webResponse)
		return
	}
	resp := c.connectorService.GetEmailFilter(req)
	ctx.JSON(resp.Code, resp)
}
