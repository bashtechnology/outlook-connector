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
// @Param body body request.GetEmailFilterRequest true "Requisição Body"
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @In header
// @Name Authorization
// @Type apiKey
// @Success 200 {object} response.HttpResponse{data=[]response.EmailResponse{}} "Dados recebidos!"
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

// GetEmailFilterFolder godoc
// @Tags Email
// @Summary Leitura de Emails
// @Description Obter os emails para os parametros informados.
// @ID GetEmailFilterFolder
// @Param body body request.GetEmailFilterRequest true "Requisição Body"
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @In header
// @Name Authorization
// @Type apiKey
// @Success 200 {object} response.HttpResponse{data=[]response.EmailResponse{}} "Dados recebidos!"
// @Failure 400 {object} response.HttpResponse "Requisição Inválida"
// @Router /v1/connector/folder/read [post]
func (c *ConnectorController) GetEmailFilterFolder(ctx *gin.Context) {
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
	resp := c.connectorService.GetEmailFilterFolder(req)
	ctx.JSON(resp.Code, resp)
}

// GetEmailFilterFull godoc
// @Tags Email
// @Summary Leitura de Emails
// @Description Obter os emails para os parametros informados.
// @ID GetEmailFilterFull
// @Param body body request.GetEmailFilterRequest true "Requisição Body"
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @In header
// @Name Authorization
// @Type apiKey
// @Success 200 {object} response.HttpResponse{} "Dados recebidos!"
// @Failure 400 {object} response.HttpResponse "Requisição Inválida"
// @Router /v1/connector/read/full [post]
func (c *ConnectorController) GetEmailFilterFull(ctx *gin.Context) {
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
	resp := c.connectorService.GetEmailFilterFull(req)
	ctx.JSON(resp.Code, resp)
}

// GetEmailFilterFullFolder godoc
// @Tags Email
// @Summary Leitura de Emails
// @Description Obter os emails para os parametros informados.
// @ID GetEmailFilterFullFolder
// @Param body body request.GetEmailFilterRequest true "Requisição Body"
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @In header
// @Name Authorization
// @Type apiKey
// @Success 200 {object} response.HttpResponse{} "Dados recebidos!"
// @Failure 400 {object} response.HttpResponse "Requisição Inválida"
// @Router /v1/connector/folder/read/full [post]
func (c *ConnectorController) GetEmailFilterFullFolder(ctx *gin.Context) {
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
	resp := c.connectorService.GetEmailFilterFullFolder(req)
	ctx.JSON(resp.Code, resp)
}

// MarkEmailID godoc
// @Tags Email
// @Summary Leitura de Emails
// @Description Obter os emails para os parametros informados.
// @ID MarkEmailID
// @Param body body request.MarkEmailIDRequest true "Requisição Body"
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @In header
// @Name Authorization
// @Type apiKey
// @Success 200 {object} response.HttpResponse{data=[]response.MarkEmailResponse{}} "Dados recebidos!"
// @Failure 400 {object} response.HttpResponse "Requisição Inválida"
// @Router /v1/connector/mark [post]
func (c *ConnectorController) MarkEmailID(ctx *gin.Context) {
	req := request.MarkEmailIDRequest{}
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
	resp := c.connectorService.MarkEmailID(req)
	ctx.JSON(resp.Code, resp)
}

// MoveTo godoc
// @Tags Email
// @Summary Leitura de Emails
// @Description Obter os emails para os parametros informados.
// @ID MoveTo
// @Param body body request.MoveToRequest true "Requisição Body"
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @In header
// @Name Authorization
// @Type apiKey
// @Success 200 {object} response.HttpResponse{data=[]response.MarkEmailResponse{}} "Dados recebidos!"
// @Failure 400 {object} response.HttpResponse "Requisição Inválida"
// @Router /v1/connector/move [post]
func (c *ConnectorController) MoveTo(ctx *gin.Context) {
	req := request.MoveToRequest{}
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
	resp := c.connectorService.MoveTo(req)
	ctx.JSON(resp.Code, resp)
}

// GetFolders godoc
// @Tags Email
// @Summary Leitura de Emails
// @Description Obter os emails para os parametros informados.
// @ID GetFolders
// @Param body body request.GetFoldersRequest true "Requisição Body"
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @In header
// @Name Authorization
// @Type apiKey
// @Success 200 {object} response.HttpResponse{data=[]response.MarkEmailResponse{}} "Dados recebidos!"
// @Failure 400 {object} response.HttpResponse "Requisição Inválida"
// @Router /v1/connector/folders [post]
func (c *ConnectorController) GetFolders(ctx *gin.Context) {
	req := request.GetFoldersRequest{}
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
	resp := c.connectorService.GetFolders(req)
	ctx.JSON(resp.Code, resp)
}
