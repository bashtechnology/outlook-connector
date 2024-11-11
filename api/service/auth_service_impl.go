package service

import (
	"outlook-connector/api/data/request"
	"outlook-connector/api/data/response"
	"outlook-connector/config"
)

type AuthServiceImpl struct {
	env config.Config
}

func NewAuthServiceImpl(env config.Config) (*AuthServiceImpl, error) {
	return &AuthServiceImpl{
		env: env,
	}, nil
}

func (s *AuthServiceImpl) GetToken(req request.GetTokenRequest) response.HttpResponse {
	resp := response.HttpResponse{}

	return resp
}
