package credentials

import "github.com/dgrijalva/jwt-go"

type Username string

var secret = []byte("_chatapp_")

func (u *Username) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": string(*u),
	})

	return token.SignedString(secret)
}
