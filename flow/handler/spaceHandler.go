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
	var space entities.Space
	if err := json.NewDecoder(r.Body).Decode(&space); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&space); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(space)
}
