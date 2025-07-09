package http

import (
	"encoding/json"
	"fmt"
	"net/http"

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
}

func setCurrent(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("ID")
	id, err := utils.IDtoInt(idStr)
	if err != nil {
		utils.HTTPErrorHandling(w, r, http.StatusBadRequest, err.Error())
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

	updatedState, err := db.ViewColor(id)
	if err != nil {
		utils.HTTPErrorHandling(w, r, http.StatusInternalServerError, "Failed to fetch lamp state")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedState)
}
