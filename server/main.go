package main

import (
	"github.com/daycolor/chat-app/user"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/register", user.Register).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
