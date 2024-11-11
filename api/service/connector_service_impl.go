package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"outlook-connector/api/data/common"
	"outlook-connector/api/data/request"
	"outlook-connector/api/data/response"
	"outlook-connector/config"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	graphusers "github.com/microsoftgraph/msgraph-sdk-go/users"
)

type ConnectorServiceImpl struct {
	mutexRotina sync.Mutex
	rotina      *time.Ticker
	env         config.Config
	client      *msgraphsdk.GraphServiceClient
}

func NewConnectorServiceImpl(env config.Config) (*ConnectorServiceImpl, error) {
	// Configurar as credenciais usando azidentity diretamente
	cred, err := azidentity.NewClientSecretCredential(env.TenantID, env.ClientID, env.ClientSecret, nil)
	if err != nil {
		return nil, fmt.Errorf("Erro ao criar as credenciais: %v", err)
	}

	// Criar o cliente da Microsoft Graph diretamente
	client, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{"https://graph.microsoft.com/.default"})
	if err != nil {
		return nil, fmt.Errorf("Erro ao criar o cliente da Graph API: %v", err)
	}

	return &ConnectorServiceImpl{
		env:    env,
		client: client,
	}, nil
}

func (s *ConnectorServiceImpl) GetEmailFilter(req request.GetEmailFilterRequest) response.HttpResponse {
	emails, err := s.LerEmails(req)
	if err != nil {
		return response.HttpResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Erro",
			Message: err.Error(),
		}
	}
	return response.HttpResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "E-mails lidos",
		Data:    emails,
	}
}

// Função para remover conteúdo anterior em uma resposta de e-mail
func extractLatestReply(emailContent string) string {
	// Padrão para detectar quando começa o conteúdo anterior de uma mensagem
	// Exemplo de padrões: "Em <data> escreveu:", ou "<nome> escreveu:"
	re := regexp.MustCompile(`(?i)em .* escreveu:`)

	// Divide o conteúdo do e-mail onde o padrão é encontrado
	parts := re.Split(emailContent, -1)

	// Retorna apenas a parte mais recente, ou seja, a primeira parte antes de qualquer "Em <data> escreveu"
	if len(parts) > 0 {
		return strings.TrimSpace(parts[0])
	}

	return emailContent // Se não encontrar o padrão, retorna o conteúdo completo
}
func getAttachments(graphClient *msgraphsdk.GraphServiceClient, messageId string, userId string) ([]common.Anexo, error) {
	attachmentsRequest := graphClient.Users().ByUserId(userId).Messages().ByMessageId(messageId).Attachments()

	attachmentsCollection, err := attachmentsRequest.Get(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("Erro ao obter anexos: %v", err)
	}

	var anexos []common.Anexo

	// Processa cada anexo
	for _, attachment := range attachmentsCollection.GetValue() {
		// Verifica se o anexo é um arquivo
		fileAttachment, ok := attachment.(models.FileAttachmentable)
		if !ok {
			continue // Ignora anexos que não são arquivos
		}

		// Codifica o conteúdo do anexo em base64
		encodedContent := base64.StdEncoding.EncodeToString(fileAttachment.GetContentBytes())

		// Adiciona o anexo ao array de anexos
		anexos = append(anexos, common.Anexo{
			Base64:      encodedContent,
			ContentType: *fileAttachment.GetContentType(),
			ID:          *fileAttachment.GetId(),
			MimeType:    *fileAttachment.GetContentType(),
			Name:        *fileAttachment.GetName(),
			URL:         "", // Preencher se houver URL
		})
	}

	return anexos, nil
}

// Função para remover tags HTML do conteúdo
func stripHTMLTags(htmlContent string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	plainText := re.ReplaceAllString(htmlContent, "")
	return strings.TrimSpace(plainText)
}

// Função para marcar um e-mail como lido
func (s *ConnectorServiceImpl) MarkEmailAsRead(emailID string) error {
	// Cria a requisição para marcar o e-mail como lido
	body := models.NewMessage()
	isRead := true
	body.SetIsRead(&isRead)

	// Atualiza o e-mail especificado pelo ID
	_, err := s.client.Users().ByUserId("webbot@bashtechnology.com.br").Messages().ByMessageId(emailID).Patch(context.Background(), body, nil)
	if err != nil {
		return fmt.Errorf("erro ao marcar e-mail como lido: %v", err)
	}
	return nil
}

// Função para ler e-mails usando a Microsoft Graph API
func (s *ConnectorServiceImpl) LerEmails(req request.GetEmailFilterRequest) ([]models.Messageable, error) {
	// Ler os e-mails da caixa de entrada
	// filter := "isRead eq false"
	// requestParameters := &graphusers.ItemMessagesRequestBuilderGetQueryParameters{
	// 	Select: []string{"sender", "subject", "isRead", "body", "conversationid"},
	// 	Filter: &filter,
	// }
	requestParameters := &graphusers.ItemMessagesRequestBuilderGetQueryParameters{
		Select:  req.Select, //[]string{"sender", "subject", "isRead", "body", "conversationid"}
		Filter:  req.Filter, //"isRead eq false"
		Search:  req.Search,
		Expand:  req.Expand,
		Orderby: req.Orderby,
		Top:     req.Top,
		Skip:    req.Skip,
	}

	configuration := &graphusers.ItemMessagesRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}
	emails, err := s.client.Users().ByUserId(s.env.MailBox).Messages().Get(context.Background(), configuration)
	if err != nil {
		return nil, fmt.Errorf("Erro ao ler e-mails: %v", err)
	}
	return emails.GetValue(), nil
}
