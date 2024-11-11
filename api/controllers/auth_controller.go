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

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(env config.Config) *AuthController {
	authService, err_init := service.NewAuthServiceImpl(env)
	if err_init != nil {
		slog.Error("⚙️AuthController", "💥Erro ao iniciar Service.", "Error", err_init)
	}
	return &AuthController{authService: authService}
}

// GetToken godoc
// @Tags Auth
// @Summary Get Token
// @Description Obter os dados de token JWT para acesso.
// @ID GetToken
// @Param body body request.GetTokenRequest true "Requisição Body"
// @Produce json
// @Security ApiKeyAuth
// @SecurityDefinitions ApiKeyAuth
// @In header
// @Name Authorization
// @Type apiKey
// @Success 200 {object} response.HttpResponse{} "Dados recebidos!"
// @Failure 400 {object} response.HttpResponse "Requisição Inválida"
// @Router /v1/open/token [post]
func (c *AuthController) GetToken(ctx *gin.Context) {
	req := request.GetTokenRequest{}
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
	resp := c.authService.GetToken(req)
	ctx.JSON(resp.Code, resp)
}
