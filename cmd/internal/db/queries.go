package db

import (
	"errors"

	"github.com/siestaw/laterna/server/cmd/internal/logger"
	"github.com/siestaw/laterna/server/cmd/internal/models"
	"github.com/siestaw/laterna/server/cmd/utils"
)

func CreateController(target int) (string, int, error) {
	token, _ := utils.GenerateToken()
	hash, _ := utils.HashToken(token)

	var maxID int
	stmt := DB.QueryRow("SELECT COALESCE(MAX(id), 0) FROM controllers")
	stmt.Scan(&maxID)
	id := maxID + 1

	_, err := DB.Exec("INSERT INTO controllers VALUES (?, ?, '#FFFFFF', CURRENT_TIMESTAMP)", id, hash)
	if err != nil {
		logger.DBLogger.Printf("Failed to create a new controller: %v", err)
		return "", id, err
	}

	_, err = DB.Exec("INSERT INTO permissions VALUES (?, ?) ", id, target)
	if err != nil {
		logger.DBLogger.Printf("Failed to set permissions for %d: %v", id, err)
		return token, id, err
	}
	return token, id, nil
}

func CreateAdmin() string {
	token, _ := utils.GenerateToken()
	hash, _ := utils.HashToken(token)

	_, err := DB.Exec("INSERT INTO controllers (id, token_hash) VALUES (0, ?)", hash)
	if err != nil {
		logger.DBLogger.Fatalf("Error creating admin user: %s", err)
	}

	return token
}

func ResetAdmin() {
	_, err := DB.Exec("DELETE FROM controllers WHERE id = 0")
	if err != nil {
		logger.DBLogger.Printf("An error occured while resetting the admin account: %s", err)
	}
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
	row := DB.QueryRow("SELECT * FROM lamp_state WHERE id = ?", id)

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

	_, err = DB.Exec("UPDATE lamp_state SET color = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?", color, id)
	if err != nil {
		logger.DBLogger.Printf("Failed to update lamp %d: %v", id, err)
		return err
	}

	logger.DBLogger.Printf("Lamp %d color updated to %s", id, color)
	return nil
}
