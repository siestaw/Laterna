package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/siestaw/laterna/server/cmd/internal/config"
	"github.com/siestaw/laterna/server/cmd/internal/db"
	"github.com/siestaw/laterna/server/cmd/internal/logger"
	"github.com/siestaw/laterna/server/cmd/internal/models"
	"github.com/siestaw/laterna/server/cmd/utils"
)

func StartHTTPServer() {
	router := http.NewServeMux()
	router.HandleFunc("GET /api/v1/id/{ID}", getCurrent)
	router.HandleFunc("PUT /api/v1/id/{ID}", setCurrent)

	if !db.ControllerExists(0) {
		adminToken := db.CreateAdmin()
		fmt.Println("IMPORTANT")
		fmt.Println("- - - - - - - - - - - - - - - - - - ")
		fmt.Println("ADMIN TOKEN:")
		fmt.Printf("%s\n", adminToken)
		fmt.Println("REGENERATE THE TOKEN BY RUNNING WITH -resetAdminToken")
		fmt.Println("- - - - - - - - - - - - - - - - - - ")
	}

	port := config.AppConfig.HTTP.Port
	logger.HTTPLogger.Printf("HTTP Server running on :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func getCurrent(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("ID")
	id, err := utils.IDtoInt(idStr)
	if err != nil {
		utils.HTTPErrorHandling(w, r, http.StatusBadRequest, err.Error())
		return
	}

	state, err := db.ViewColor(id)
	if err != nil {
		utils.HTTPErrorHandling(w, r, http.StatusBadRequest, "Lamp not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(state)
	defer r.Body.Close()
}

func setCurrent(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("ID")
	id, err := utils.IDtoInt(idStr)
	if err != nil {
		utils.HTTPErrorHandling(w, r, http.StatusBadRequest, err.Error())
		return
	}
	currentState, err := db.ViewColor(id)
	if err != nil {
		utils.HTTPErrorHandling(w, r, http.StatusInternalServerError, "Failed to get current lamp color")
		return
	}

	if time.Since(currentState.UpdatedAt).Seconds() < config.AppConfig.HTTP.Cooldown {
		utils.HTTPErrorHandling(w, r, http.StatusTooManyRequests, "Slow down!")
		return
	}

	var req models.LampUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.HTTPLogger.Print(err)
		utils.HTTPErrorHandling(w, r, http.StatusBadRequest, "Invalid JSON")
		return
	}

	err = db.SetColor(id, req.Color)
	if err != nil {
		logger.HTTPLogger.Printf("Could not update lamp %d: %s", id, err)
		utils.HTTPErrorHandling(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	updatedState := currentState
	updatedState.Color = req.Color
	updatedState.UpdatedAt = time.Now()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedState)
	defer r.Body.Close()
}
