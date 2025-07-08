package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/siestaw/laterna/server/cmd/internal/models"
)

func HTTPErrorHandling(w http.ResponseWriter, r *http.Request, status int, message string) {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	errResp := models.HTTPError{
		Timestamp: timestamp,
		Status:    status,
		Error:     http.StatusText(status),
		Message:   message,
		Path:      r.URL.Path,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errResp)
}
