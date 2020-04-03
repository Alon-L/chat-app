package user

import (
	"context"
	"encoding/json"
	"github.com/daycolor/chat-app/credentials"
	"github.com/daycolor/chat-app/mongo"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cred credentials.Credentials

	err := json.NewDecoder(r.Body).Decode(&cred)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	password := cred.Password.Salt()
	_, err = cred.Password.Hash()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if match := cred.Password.Compare(password); !match {
		http.Error(w, "Password hashing failed", http.StatusInternalServerError)
		return
	}

	database, err := mongo.ConnectDB()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token, err := cred.Username.GenerateToken()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	users := database.Collection("users")

	res, err := users.InsertOne(context.TODO(), mongo.User{
		Credentials: &cred,
		Token:       token,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(res)
}
