package db

import (
	"github.com/siestaw/laterna/server/cmd/internal/logger"
	"github.com/siestaw/laterna/server/cmd/internal/models"
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
