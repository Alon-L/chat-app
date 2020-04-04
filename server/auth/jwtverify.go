package auth

import (
	"context"
	"github.com/daycolor/chat-app/models"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("x-access-token")
		token := strings.TrimSpace(header)

		if token == "" {
			http.Error(w, "Missing auth token", http.StatusForbidden)
			return
		}

		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(token, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), "token", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
