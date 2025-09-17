package handler

import (
	"anytype-flow-crud/flow/entities"
	"anytype-flow-crud/flow/service"
	"encoding/json"
	"net/http"
)

type SpaceHandler struct {
	service *service.SpaceService
}

func NewSpaceHandler(s *service.SpaceService) *SpaceHandler {
	return &SpaceHandler{service: s}
}

func (h *SpaceHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req entities.Space
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
