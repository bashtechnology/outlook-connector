package common

type ParamsFluxo struct {
	FluxoInicial string
}

type ParamsSender struct {
	Sandbox       bool
	AppName       string
	Token         string
	Number        string
	To            string
	Body          string
	PlataformasID int
}

type Anexo struct {
	Base64      string `json:"base64"`
	ContentType string `json:"content_type"`
	ID          string `json:"id"`
	MimeType    string `json:"mime_type"`
	Name        string `json:"name"`
	URL         string `json:"url"`
}

type Payload struct {
	Anexos  []Anexo `json:"anexos"`
	Caption string  `json:"caption"`
	// History []string `json:"history"`
	ID         string `json:"id"`
	IDConversa string `json:"id_conversa"`
	Sender     string `json:"sender"`
	Text       string `json:"text"`
	Html       string `json:"html"`
}

type EmailData struct {
	App       string  `json:"app"`
	Payload   Payload `json:"payload"`
	Timestamp int64   `json:"timestamp"`
	Type      string  `json:"type"`
	Version   int     `json:"version"`
}
