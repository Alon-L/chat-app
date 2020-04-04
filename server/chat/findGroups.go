package chat

import (
	"context"
	"encoding/json"
	"github.com/daycolor/chat-app/models"
	"github.com/daycolor/chat-app/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func FindGroups(w http.ResponseWriter, r *http.Request) {
	token := r.Context().Value("token").(*models.Token)

	database, err := mongo.ConnectDB()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	groups := database.Collection("groups")

	cur, err := groups.Find(context.TODO(), bson.M{"participants": token.ID})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	results := &[]*models.Group{}

	err = cur.All(context.TODO(), results)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(results)
}
