package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/siestaw/laterna/server/cmd/internal/config"
	"github.com/siestaw/laterna/server/cmd/internal/db"
	"github.com/siestaw/laterna/server/cmd/internal/logger"
	"github.com/siestaw/laterna/server/cmd/internal/middleware"
	"github.com/siestaw/laterna/server/cmd/internal/models"
	"github.com/siestaw/laterna/server/cmd/utils"
)

func RegisterColorRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/colors/", middleware.WithAdminAuth(listColors))
	mux.HandleFunc("GET /api/v1/colors/{ID}", middleware.WithAdminAuth(getCurrent))
	mux.HandleFunc("PUT /api/v1/colors/{ID}", middleware.WithAdminAuth(setCurrent))
}

func listColors(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	colors, err := db.GetAllColors()
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	utils.SuccessResponse(w, http.StatusOK, colors)
}

func getCurrent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idStr := r.PathValue("ID")
	id, err := utils.IDtoInt(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if !db.ControllerExists(id) {
		utils.ErrorResponse(w, http.StatusNotFound, "Lamp does not exist")
		return
	}

	state, err := db.ViewColor(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(w, http.StatusOK, state)
}

func setCurrent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idStr := r.PathValue("ID")
	id, err := utils.IDtoInt(idStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if !db.ControllerExists(id) {
		utils.ErrorResponse(w, http.StatusBadRequest, "Lamp does not exist")
		return
	}

	currentState, err := db.ViewColor(id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "Failed to get current lamp color")
		return
	}

	if time.Since(currentState.UpdatedAt).Seconds() < config.AppConfig.HTTP.Cooldown {
		utils.ErrorResponse(w, http.StatusTooManyRequests, "Slow down!")
		return
	}

	var req models.LampUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.HTTPLogger.Print(err)
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	err = db.SetColor(id, req.Color)
	if err != nil {
		logger.HTTPLogger.Printf("Could not update lamp %d: %s", id, err)
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	updatedState := currentState
	updatedState.Color = req.Color
	updatedState.UpdatedAt = time.Now()
	utils.SuccessResponse(w, http.StatusOK, updatedState)

}
