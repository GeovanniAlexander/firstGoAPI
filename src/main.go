package main

import (
	"fmt"
	"net/http"

	"github.com/GeovanniAlexander/01-firtsGoAPI/src/api"
	"github.com/gorilla/mux"
)

func main() {
	var port string = "8080"

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/").Subrouter()
	apiRouter.HandleFunc("/todos", api.CreateTodo).Methods(http.MethodPost)
	apiRouter.HandleFunc("/todos/{id}", api.GetTodo).Methods(http.MethodGet)

	fmt.Printf("Running in port %s", port)
	http.ListenAndServe(":"+port, router)

}
