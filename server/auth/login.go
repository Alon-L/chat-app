package auth

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/daycolor/chat-app/models"
	"github.com/daycolor/chat-app/mongo"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

var secret = "_chatapp_"

func findUser(username string, password models.Password) (string, error) {
	user := &models.User{}

	database, err := mongo.ConnectDB()
	if err != nil {
		return "", err
	}

	users := database.Collection("users")

	err = users.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return "", err
	}

	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	password.Salt()

	if match := user.Password.Compare(password); !match {
		return "", errors.New("invalid password")
	}

	tk := models.Token{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode credentials
	user := &models.User{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := findUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(w).Encode(res)
}
