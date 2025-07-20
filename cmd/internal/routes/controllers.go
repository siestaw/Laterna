package routes

import (
	"encoding/json"
	"net/http"

	"github.com/siestaw/laterna/server/cmd/internal/db"
	"github.com/siestaw/laterna/server/cmd/internal/logger"
	"github.com/siestaw/laterna/server/cmd/internal/middleware"
	"github.com/siestaw/laterna/server/cmd/internal/models"
	"github.com/siestaw/laterna/server/cmd/utils"
)

func RegisterControllerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/controllers", middleware.WithAdminAuth(createController))
	mux.HandleFunc("DELETE /api/v1/controllers", middleware.WithAdminAuth(deleteController))
}

func createController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ID, err := db.CreateController()
	if err != nil {
		utils.HTTPResponseHandler(w, r, http.StatusInternalServerError, "An error occured. Check the server logs for more information")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int64{"ID": ID})
}

func deleteController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req models.ControllerRequests
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		logger.HTTPLogger.Print(err)
		utils.HTTPResponseHandler(w, r, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if !db.ControllerExists(req.ID) || req.ID <= 0 {
		utils.HTTPResponseHandler(w, r, http.StatusBadRequest,"Invalid ID")
		return
	}
	if db.DeleteController(req.ID) != nil {
		utils.HTTPResponseHandler(w, r, http.StatusInternalServerError,"An error occured. Check the server logs for more information")
		return
	}
	utils.HTTPResponseHandler(w, r, http.StatusOK, "Success")
}