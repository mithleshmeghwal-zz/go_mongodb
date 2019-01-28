package main

import (
	"log"
	"net/http"

	"go_mongodb_mux/musicstore/album"

	"github.com/gorilla/handlers"
)

func main() {
	router := album.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
