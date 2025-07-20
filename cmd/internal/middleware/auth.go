package middleware

import (
	"net/http"

	"github.com/siestaw/laterna/server/cmd/internal/db"
	"github.com/siestaw/laterna/server/cmd/utils"
)

func WithAdminAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !db.IsAdmin(token) {
			utils.HTTPResponseHandler(w, r, http.StatusUnauthorized, "Invalid token")
			return 
		}
		handlerFunc(w, r)
	}
}