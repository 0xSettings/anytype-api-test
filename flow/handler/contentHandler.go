package handler

import "anytype-flow-crud/flow/service"

type ContentHandler struct {
	service *service.ContentService
}

func NewContentHandler(ser *service.ContentService) *ContentHandler {
	return &ContentHandler{
		service: ser,
	}
}
