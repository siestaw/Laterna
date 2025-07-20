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
	mux.HandleFunc("GET /api/v1/colors/{ID}", middleware.WithAdminAuth(getCurrent))
	mux.HandleFunc("PUT /api/v1/colors/{ID}", middleware.WithAdminAuth(setCurrent))
//	mux.HandleFunc("WS /api/v1/ws/colors/{ID}", setCurrentWebsocket)
}

func getCurrent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idStr := r.PathValue("ID")
	id, err := utils.IDtoInt(idStr)
	if err != nil {
		utils.HTTPResponseHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	
	if !db.ControllerExists(id) {
		utils.HTTPResponseHandler(w, r, http.StatusBadRequest, "Lamp does not exist")
		return
	}

	state, err := db.ViewColor(id)
	if err != nil {
		utils.HTTPResponseHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(state)
}

func setCurrent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	idStr := r.PathValue("ID")
	id, err := utils.IDtoInt(idStr)
	if err != nil {
		utils.HTTPResponseHandler(w, r, http.StatusBadRequest, err.Error())
		return
	}
	if !db.ControllerExists(id) {
		utils.HTTPResponseHandler(w, r, http.StatusBadRequest, "Lamp does not exist")
		return
	}

	currentState, err := db.ViewColor(id)
	if err != nil {
		utils.HTTPResponseHandler(w, r, http.StatusInternalServerError, "Failed to get current lamp color")
		return
	}

	if time.Since(currentState.UpdatedAt).Seconds() < config.AppConfig.HTTP.Cooldown {
		utils.HTTPResponseHandler(w, r, http.StatusTooManyRequests, "Slow down!")
		return
	}

	var req models.LampUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.HTTPLogger.Print(err)
		utils.HTTPResponseHandler(w, r, http.StatusBadRequest, "Invalid JSON")
		return
	}

	err = db.SetColor(id, req.Color)
	if err != nil {
		logger.HTTPLogger.Printf("Could not update lamp %d: %s", id, err)
		utils.HTTPResponseHandler(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	updatedState := currentState
	updatedState.Color = req.Color
	updatedState.UpdatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedState)
	logger.HTTPLogger.Printf("Lamp %d updated to color %s", id, req.Color)
}