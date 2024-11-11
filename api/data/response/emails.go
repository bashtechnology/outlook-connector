package response

import "time"

type EmailResponse struct {
	ID                *string              `json:"id,omitempty"`
	Subject           *string              `json:"subject,omitempty"`
	BodyContent       *string              `json:"body_content,omitempty"`
	BodyContentType   *string              `json:"body_content_type,omitempty"`
	SenderName        *string              `json:"sender_name,omitempty"`
	SenderEmail       *string              `json:"sender_email,omitempty"`
	FromName          *string              `json:"from_name,omitempty"`
	FromEmail         *string              `json:"from_email,omitempty"`
	ToRecipients      []RecipientResponse  `json:"to_recipients,omitempty"`
	CcRecipients      []RecipientResponse  `json:"cc_recipients,omitempty"`
	BccRecipients     []RecipientResponse  `json:"bcc_recipients,omitempty"`
	ReceivedDateTime  *time.Time           `json:"received_date_time,omitempty"`
	SentDateTime      *time.Time           `json:"sent_date_time,omitempty"`
	IsRead            *bool                `json:"is_read,omitempty"`
	HasAttachments    *bool                `json:"has_attachments,omitempty"`
	ConversationID    *string              `json:"conversation_id,omitempty"`
	InternetMessageID *string              `json:"internet_message_id,omitempty"`
	Importance        *string              `json:"importance,omitempty"`
	Attachments       []AttachmentResponse `json:"attachments,omitempty"`
}

type RecipientResponse struct {
	Name    *string `json:"name,omitempty"`
	Address *string `json:"address,omitempty"`
}

type AttachmentResponse struct {
	Base64               string     `json:"base64"`
	Name                 *string    `json:"name,omitempty"`                    // Nome do anexo
	ContentType          *string    `json:"content_type,omitempty"`            // Tipo MIME do anexo (ex: image/png, application/pdf)
	Size                 *int32     `json:"size,omitempty"`                    // Tamanho do anexo em bytes
	IsInline             *bool      `json:"is_inline,omitempty"`               // Indica se o anexo é inline
	LastModifiedDateTime *time.Time `json:"last_modified_date_time,omitempty"` // Data e hora da última modificação do anexo
	Id                   *string    `json:"id,omitempty"`                      // ID do anexo (caso disponível)
}
