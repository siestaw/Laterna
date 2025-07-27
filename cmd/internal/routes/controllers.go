package routes

import (
	"encoding/json"
	"net/http"

	"github.com/siestaw/laterna/server/cmd/internal/db"
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
		utils.ErrorResponse(w, http.StatusInternalServerError, "An error occured. Check the server logs for more information")
	}
	utils.SuccessResponse(w, http.StatusOK, models.CreateData{Created: int(ID)})
}

func deleteController(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req models.ControllerRequests
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid JSON")
		return
	}
	if !db.ControllerExists(req.ID) || req.ID <= 0 {
		utils.ErrorResponse(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	if db.DeleteController(req.ID) != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "An error occured. Check the server logs for more information")
		return
	}
	utils.SuccessResponse(w, http.StatusOK, models.DeleteData{Deleted: req.ID})
}
