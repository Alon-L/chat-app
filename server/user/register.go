package user

import (
	"context"
	"encoding/json"
	"github.com/daycolor/chat-app/mongo"
	"net/http"
)

func Register(w http.ResponseWriter, _ *http.Request, registration Registration) {
	w.Header().Set("Content-Type", "application/json")

	password := registration.Password.Salt()
	_, err := registration.Password.Hash()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if match := registration.Password.Compare(password); !match {
		http.Error(w, "Password hashing failed", http.StatusInternalServerError)
		return
	}

	database, err := mongo.ConnectDB()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := database.Collection("users")
	res, err := users.InsertOne(context.TODO(), registration)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(res)
}
