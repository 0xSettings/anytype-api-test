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
	var req entities.Page
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created, err := h.service.Create(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}
