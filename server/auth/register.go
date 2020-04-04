package auth

import (
	"context"
	"encoding/json"
	"github.com/daycolor/chat-app/models"
	"github.com/daycolor/chat-app/mongo"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode credentials
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Salt and hash the password
	user.Password.Salt()
	err = user.Password.Hash()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Connect to DB and insert data
	database, err := mongo.ConnectDB()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := database.Collection("users")

	res, err := users.InsertOne(context.TODO(), user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(res)
}
