package utils

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/siestaw/laterna/server/cmd/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func IDtoInt(id string) (int, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("invalid ID format")
	}
	if idInt <= 0 {
		return 0, errors.New("invalid ID")
	}
	return idInt, nil
}

func IsValidHexColor(color string) bool {
	pattern := `^#?([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$`
	match, _ := regexp.MatchString(pattern, color)
	return match
}

func GenerateToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func HashToken(token string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	return string(hash), err
}

func ValidateToken(providedToken string, storedHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(providedToken))
	return err == nil
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, res models.HTTPResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	res.Status = statusCode
	res.Timestamp = time.Now().UTC().Format(time.RFC3339)

	json.NewEncoder(w).Encode(res)
}

func SuccessResponse(w http.ResponseWriter, statusCode int, data any) {
	WriteJSONResponse(w, statusCode, models.HTTPResponse{
		Success: true,
		Data:    data,
	})
}

func ErrorResponse(w http.ResponseWriter, statusCode int, err string) {
	WriteJSONResponse(w, statusCode, models.HTTPResponse{
		Success: false,
		Error:   err,
	})
}
