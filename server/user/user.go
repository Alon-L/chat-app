package user

import (
	"encoding/json"
	"github.com/daycolor/chat-app/credentials"
	"net/http"
)

type Registration struct {
	Username string
	Password credentials.Password
}

type Route func(http.ResponseWriter, *http.Request, Registration)

func Middleware(next Route) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var registration Registration

		err := json.NewDecoder(r.Body).Decode(&registration)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		next(w, r, registration)
	}
}
