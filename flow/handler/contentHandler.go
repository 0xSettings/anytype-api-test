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
	var content entities.Content

	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := hand.service.Create(&content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(content)
}

func (hand *ContentHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var content entities.Content
	if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := hand.service.Update(id, &content); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(content)
}

func (hand *ContentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := hand.service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
