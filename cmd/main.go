package main

import (
	"anytype-flow-crud/flow/handler"
	"anytype-flow-crud/flow/repository"
	"anytype-flow-crud/flow/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	repo := repository.NewFlowRepo()

	spaceService := service.NewSpaceService(repo)
	pageService := service.NewPageService(repo)
	contentService := service.NewContentService(repo)

	spaceHandler := handler.NewSpaceHandler(spaceService)
	pageHandler := handler.NewPageHandler(pageService)
	contentHandler := handler.NewContentHandler(contentService)

	// API routes
	r.HandleFunc("/api/v1/spaces", spaceHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/pages", pageHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/content", contentHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/content/{id}", contentHandler.Update).Methods("PUT")
	r.HandleFunc("/api/v1/content/{id}", contentHandler.Delete).Methods("DELETE")

	log.Println("Server Running on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
