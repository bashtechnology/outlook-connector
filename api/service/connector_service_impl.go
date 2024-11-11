package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
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
	"github.com/k0kubun/pp"
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
	emailList := []*response.EmailResponse{}

	for _, email := range emails {
		item, err := s.ConvertEmailToStruct(email)
		if err != nil {
			log.Printf("Erro ao processar email: %v", err)
			continue
		}
		emailList = append(emailList, item)
	}
	return response.HttpResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "E-mails lidos",
		Data:    emailList,
	}
}
func (s *ConnectorServiceImpl) GetEmailFilterFull(req request.GetEmailFilterRequest) response.HttpResponse {
	emails, err := s.LerEmails(req)
	if err != nil {
		return response.HttpResponse{
			Code:    http.StatusInternalServerError,
			Status:  "Erro",
			Message: err.Error(),
		}
	}
	pp.ColoringEnabled = false // Desativa as cores
	emailList := []json.RawMessage{}
	for _, email := range emails {
		item := pp.Sprintln(email)
		emailJSON, err := json.Marshal(item)
		if err != nil {
			return response.HttpResponse{
				Code:    http.StatusInternalServerError,
				Status:  "Erro",
				Message: "Erro ao converter email para JSON: " + err.Error(),
			}
		}
		emailList = append(emailList, emailJSON)
	}
	return response.HttpResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "E-mails lidos",
		Data:    emailList,
	}
}

func (s *ConnectorServiceImpl) ConvertEmailToStruct(email models.Messageable) (*response.EmailResponse, error) {
	resp := &response.EmailResponse{}
	if id := email.GetId(); id != nil {
		resp.ID = id
	}
	if subject := email.GetSubject(); subject != nil {
		resp.Subject = subject
	}
	if body := email.GetBody(); body != nil {
		resp.BodyContent = body.GetContent()
		contentType := body.GetContentType().String()
		resp.BodyContentType = &contentType
	}
	if sender := email.GetSender(); sender != nil && sender.GetEmailAddress() != nil {
		resp.SenderName = sender.GetEmailAddress().GetName()
		resp.SenderEmail = sender.GetEmailAddress().GetAddress()
	}
	if from := email.GetFrom(); from != nil && from.GetEmailAddress() != nil {
		resp.FromName = from.GetEmailAddress().GetName()
		resp.FromEmail = from.GetEmailAddress().GetAddress()
	}
	if toRecipients := email.GetToRecipients(); toRecipients != nil {
		for _, recipient := range toRecipients {
			resp.ToRecipients = append(resp.ToRecipients, response.RecipientResponse{
				Name:    recipient.GetEmailAddress().GetName(),
				Address: recipient.GetEmailAddress().GetAddress(),
			})
		}
	}
	if ccRecipients := email.GetCcRecipients(); ccRecipients != nil {
		for _, recipient := range ccRecipients {
			resp.CcRecipients = append(resp.CcRecipients, response.RecipientResponse{
				Name:    recipient.GetEmailAddress().GetName(),
				Address: recipient.GetEmailAddress().GetAddress(),
			})
		}
	}
	if bccRecipients := email.GetBccRecipients(); bccRecipients != nil {
		for _, recipient := range bccRecipients {
			resp.BccRecipients = append(resp.BccRecipients, response.RecipientResponse{
				Name:    recipient.GetEmailAddress().GetName(),
				Address: recipient.GetEmailAddress().GetAddress(),
			})
		}
	}
	if receivedDateTime := email.GetReceivedDateTime(); receivedDateTime != nil {
		resp.ReceivedDateTime = receivedDateTime
	}
	if sentDateTime := email.GetSentDateTime(); sentDateTime != nil {
		resp.SentDateTime = sentDateTime
	}
	if isRead := email.GetIsRead(); isRead != nil {
		resp.IsRead = isRead
	}
	if hasAttachments := email.GetHasAttachments(); hasAttachments != nil {
		resp.HasAttachments = hasAttachments
	}
	if conversationID := email.GetConversationId(); conversationID != nil {
		resp.ConversationID = conversationID
	}
	if internetMessageID := email.GetInternetMessageId(); internetMessageID != nil {
		resp.InternetMessageID = internetMessageID
	}
	if importance := email.GetImportance(); importance != nil {
		importanceStr := importance.String()
		resp.Importance = &importanceStr
	}
	if hasAttachments := email.GetHasAttachments(); hasAttachments != nil && *hasAttachments {
		resp.HasAttachments = hasAttachments
		if attachmentList, err := s.client.Users().ByUserId(s.env.MailBox).Messages().ByMessageId(*email.GetId()).Attachments().Get(context.Background(), nil); err == nil {
			for _, attachment := range attachmentList.GetValue() {
				fileAttachment, ok := attachment.(models.FileAttachmentable)
				if ok {
					encodedContent := base64.StdEncoding.EncodeToString(fileAttachment.GetContentBytes())
					attachmentResp := response.AttachmentResponse{
						Base64:               encodedContent,
						Name:                 fileAttachment.GetName(),
						ContentType:          fileAttachment.GetContentType(),
						Size:                 fileAttachment.GetSize(),
						IsInline:             fileAttachment.GetIsInline(),
						LastModifiedDateTime: fileAttachment.GetLastModifiedDateTime(),
						Id:                   fileAttachment.GetId(),
					}
					resp.Attachments = append(resp.Attachments, attachmentResp)
				}
			}
		}
	}
	return resp, nil
}

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
func stripHTMLTags(htmlContent string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	plainText := re.ReplaceAllString(htmlContent, "")
	return strings.TrimSpace(plainText)
}
func (s *ConnectorServiceImpl) MarkEmailAsRead(emailID string) error {
	// Cria a requisição para marcar o e-mail como lido
	body := models.NewMessage()
	isRead := true
	body.SetIsRead(&isRead)

	// Atualiza o e-mail especificado pelo ID
	_, err := s.client.Users().ByUserId(s.env.MailBox).Messages().ByMessageId(emailID).Patch(context.Background(), body, nil)
	if err != nil {
		return fmt.Errorf("erro ao marcar e-mail como lido: %v", err)
	}
	return nil
}
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
