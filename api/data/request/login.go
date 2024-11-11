package request

type GetTokenRequest struct {
	Key string `validate:"required" json:"key"`
}
