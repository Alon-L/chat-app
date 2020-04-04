package main

import (
	"github.com/daycolor/chat-app/auth"
	"github.com/daycolor/chat-app/chat"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/login", auth.Login).Methods("POST")

	s := r.PathPrefix("/auth").Subrouter()

	s.Use(auth.JwtVerify)

	s.HandleFunc("/groups/create", chat.CreateGroup).Methods("POST")
	s.HandleFunc("/groups/find", chat.FindGroups).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
