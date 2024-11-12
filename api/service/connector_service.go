package service

import (
	"outlook-connector/api/data/request"
	"outlook-connector/api/data/response"
)

type ConnectorService interface {
	GetEmailFilter(req request.GetEmailFilterRequest) response.HttpResponse
	GetEmailFilterFull(req request.GetEmailFilterRequest) response.HttpResponse
	MarkEmailID(req request.MarkEmailIDRequest) response.HttpResponse
	MoveTo(req request.MoveToRequest) response.HttpResponse
	GetFolders(req request.GetFoldersRequest) response.HttpResponse
}
