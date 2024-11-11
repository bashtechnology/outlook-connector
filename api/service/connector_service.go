package service

import (
	"outlook-connector/api/data/request"
	"outlook-connector/api/data/response"
)

type ConnectorService interface {
	GetEmailFilter(req request.GetEmailFilterRequest) response.HttpResponse
}
