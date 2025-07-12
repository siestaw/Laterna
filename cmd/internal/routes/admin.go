package routes

import (
	"net/http"

	"github.com/siestaw/laterna/server/cmd/internal/db"
	"github.com/siestaw/laterna/server/cmd/utils"
)

func RegisterAdminRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/admin/controllers", createController)
}

func createController(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if !db.IsAdmin(auth) {
		utils.HTTPErrorHandling(w, r, http.StatusUnauthorized, "Invalid token")
		return
	}

}
