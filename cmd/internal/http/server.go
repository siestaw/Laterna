package http

import (
	"fmt"
	"net/http"

	"github.com/siestaw/laterna/server/cmd/internal/config"
	"github.com/siestaw/laterna/server/cmd/internal/db"
	"github.com/siestaw/laterna/server/cmd/internal/logger"
	"github.com/siestaw/laterna/server/cmd/internal/routes"
)

func StartHTTPServer() {
	router := http.NewServeMux()
	routes.RegisterRoutes(router)

	if !db.ControllerExists(0) {
		adminToken := db.CreateController(0)
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
