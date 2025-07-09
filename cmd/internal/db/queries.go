package db

import (
	"errors"

	"github.com/siestaw/laterna/server/cmd/internal/logger"
	"github.com/siestaw/laterna/server/cmd/internal/models"
	"github.com/siestaw/laterna/server/cmd/utils"
)

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
