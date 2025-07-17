package db

import (
	"errors"

	"github.com/siestaw/laterna/server/cmd/internal/logger"
	"github.com/siestaw/laterna/server/cmd/internal/models"
	"github.com/siestaw/laterna/server/cmd/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateController(id int) string {
	token, _ := utils.GenerateToken()
	hash, _ := utils.HashToken(token)

	_, err := DB.Exec("INSERT INTO controllers (id, token_hash) VALUES (?, ?)", id, hash)
	if err != nil {
		logger.DBLogger.Fatalf("Error creating controller with ID %v: %s",id, err)
	}

	return token
}

func DeleteController(id int) error {
	_, err := DB.Exec("DELETE FROM controllers WHERE id = ?", id)
	if err != nil {
		logger.DBLogger.Printf("An error occured while deleting the controller %v: %s", id, err)
		return err
	}
	return nil
}

func IsAdmin(token string) bool {
	stmt := DB.QueryRow("SELECT token_hash FROM controllers WHERE id = 0")

	var hashToken string
	err := stmt.Scan(&hashToken)
	if err != nil {
		logger.DBLogger.Printf("Error retrieving admin hash: %v", err)
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashToken), []byte(token))
	return err == nil
}

func ControllerExists(id int) bool {
	row := DB.QueryRow("SELECT COUNT(1) FROM controllers WHERE id = ?", id)

	var count int
	err := row.Scan(&count)
	if err != nil {
		logger.DBLogger.Printf("ControllerExists: %v", err)
	}
	return count > 0
}

func ViewColor(id int) (*models.LampState, error) {
	row := DB.QueryRow("SELECT id, color, updated_at FROM controllers WHERE id = ?", id)

	var state models.LampState
	err := row.Scan(&state.ID, &state.Color, &state.UpdatedAt)
	if err != nil {
		logger.DBLogger.Print(err)
		return nil, err
	}

	return &state, nil
}

func SetColor(id int, color string) error {
	if !utils.IsValidHexColor(color) {
		return errors.New("invalid HEX Code: %s")
	}
	currentColor, err := ViewColor(id)
	if err != nil {
		return err
	}
	if currentColor.Color == color {
		logger.DBLogger.Printf("Lamp %d already has color %s - skipping update", id, color)
		return nil
	}

	_, err = DB.Exec("UPDATE controllers SET color = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?", color, id)
	if err != nil {
		logger.DBLogger.Printf("Failed to update lamp %d: %v", id, err)
		return err
	}

	logger.DBLogger.Printf("Lamp %d color updated to %s", id, color)
	return nil
}
