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

	rep := repository.NewFlowRepo()
	ser := service.ExposeNewContentService(rep)
	hand := handler.NewContentHandler(ser)

	r.HandleFunc("/api/v1/content", hand.Create).Methods("POST")
	r.HandleFunc("/api/v1/content/{id}", hand.Update).Methods("PUT")
	r.HandleFunc("/api/v1/content/{id}", hand.Delete).Methods("DELETE")

	log.Println("Server Running on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
