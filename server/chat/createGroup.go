package chat

import (
	"context"
	"encoding/json"
	"github.com/daycolor/chat-app/models"
	"github.com/daycolor/chat-app/mongo"
	"net/http"
)

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(*models.Token)

	group := &models.Group{}

	err := json.NewDecoder(r.Body).Decode(group)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	group.Participants = append(group.Participants, token.ID)

	database, err := mongo.ConnectDB()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	groups := database.Collection("groups")
	res, err := groups.InsertOne(context.TODO(), group)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(res)
}
