package handler

import (
	"anytype-flow-crud/flow/entities"
	"anytype-flow-crud/flow/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ContentHandler struct {
	service *service.ContentService
}

func NewContentHandler(ser *service.ContentService) *ContentHandler {
	return &ContentHandler{
		service: ser,
	}
}

func (hand *ContentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req entities.Content
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	created, err := hand.service.Create(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

func (hand *ContentHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var req entities.Content
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updated, err := hand.service.Update(id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (hand *ContentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := hand.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
