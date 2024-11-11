package service

import (
	"outlook-connector/api/data/request"
	"outlook-connector/api/data/response"
)

type AuthService interface {
	GetToken(req request.GetTokenRequest) response.HttpResponse
}
