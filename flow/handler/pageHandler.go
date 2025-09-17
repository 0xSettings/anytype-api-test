package handler

import (
	"anytype-flow-crud/flow/entities"
	"anytype-flow-crud/flow/service"
	"encoding/json"
	"net/http"
)

type PageHandler struct {
	service *service.PageService
}

func NewPageHandler(s *service.PageService) *PageHandler {
	return &PageHandler{service: s}
}

func (h *PageHandler) Create(w http.ResponseWriter, r *http.Request) {
	var page entities.Page
	if err := json.NewDecoder(r.Body).Decode(&page); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&page); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(page)
}
