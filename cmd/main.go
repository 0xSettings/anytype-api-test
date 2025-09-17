package main

import (
	"anytype-flow-crud/flow/handler"
	"anytype-flow-crud/flow/repository"
	"anytype-flow-crud/flow/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Optional: allow overriding Anytype base url via env
	baseURL := os.Getenv("ANYTYPE_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:31009/v1"
	}
	apiKey := os.Getenv("ANYTYPE_API_KEY") // set if needed
	repo := repository.NewFlowRepoWithConfig(baseURL, apiKey)

	spaceService := service.NewSpaceService(repo)
	pageService := service.NewPageService(repo)
	contentService := service.NewContentService(repo)

	spaceHandler := handler.NewSpaceHandler(spaceService)
	pageHandler := handler.NewPageHandler(pageService)
	contentHandler := handler.NewContentHandler(contentService)

	// API routes
	r.HandleFunc("/api/v1/spaces", spaceHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/objects", pageHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/content", contentHandler.Create).Methods("POST")
	r.HandleFunc("/api/v1/content/{id}", contentHandler.Update).Methods("PUT")
	r.HandleFunc("/api/v1/content/{id}", contentHandler.Delete).Methods("DELETE")

	log.Println("Server Running on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
