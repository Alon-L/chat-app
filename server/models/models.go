package models

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password Password           `bson:"password,omitempty" json:"password,omitempty"`
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
}

type Token struct {
	*User
	jwt.StandardClaims
}

type Group struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Participants []primitive.ObjectID `bson:"participants,omitempty" json:"participants,omitempty"`
}
