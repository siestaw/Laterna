package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/siestaw/laterna/server/cmd/internal/config"
	"github.com/siestaw/laterna/server/cmd/internal/db"
	"github.com/siestaw/laterna/server/cmd/internal/logger"
	"github.com/siestaw/laterna/server/cmd/utils"
)

func StartHTTPServer() {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/id/{ID}", getCurrent)
	router.HandleFunc("POST /api/v1/id/{ID}", setCurrent)

	port := config.AppConfig.HTTP.Port
	logger.HTTPLogger.Printf("HTTP Server running on :%v", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
}

func getCurrent(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("ID")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.HTTPErrorHandling(w, r, http.StatusBadRequest, "Invalid ID")
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

}
