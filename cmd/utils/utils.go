package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/siestaw/laterna/server/cmd/internal/models"
)

func IDtoInt(id string) (int, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("invalid ID format")
	}
	if idInt < 0 {
		return 0, errors.New("ID must not be negative")
	}
	return idInt, nil
}
func IsValidHexColor(color string) bool {
	pattern := `^#?([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`
	match, _ := regexp.MatchString(pattern, color)
	return match
}

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
