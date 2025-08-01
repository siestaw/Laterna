package http

import (
	"fmt"
	"net/http"

	"github.com/siestaw/laterna/server/cmd/internal/config"
	"github.com/siestaw/laterna/server/cmd/internal/db"
	"github.com/siestaw/laterna/server/cmd/internal/logger"
	"github.com/siestaw/laterna/server/cmd/internal/routes"
	"github.com/siestaw/laterna/server/cmd/utils"
)

func StartHTTPServer() {
	router := http.NewServeMux()
	routes.RegisterColorRoutes(router)
	routes.RegisterControllerRoutes(router)
	router.HandleFunc("/", notFoundHandler)

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
	logger.HTTPLogger.Printf("Starting HTTP Server on :%v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		logger.HTTPLogger.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.ErrorResponse(w, http.StatusNotFound, "Not found")
}
