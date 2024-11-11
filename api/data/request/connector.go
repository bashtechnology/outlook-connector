package request

type GetEmailFilterRequest struct {
	Filter string `validate:"required" json:"filter"`
}
