package http

import (
	"fmt"
	"net/http"

	"github.com/mizuw/laterna/server/cmd/internal/config"
	"github.com/mizuw/laterna/server/cmd/internal/logger"
)

func StartHTTPServer() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/admin/token/new", createToken)

	port := config.AppConfig.HTTP.Port
	logger.HTTPLogger.Printf("HTTP Server running on :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func createToken(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	adminToken := config.AppConfig.HTTP.AdminToken

	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}
	if authHeader != adminToken {
		http.Error(w, "nuh uh", http.StatusUnauthorized)
		return
	}
	print(authHeader)
	fmt.Fprintf(w, "Success!!! Token: %v, authHeader: %v", adminToken, authHeader)
}
