package main

import (
	"github.com/mizuw/laterna/server/cmd/internal/config"
	"github.com/mizuw/laterna/server/cmd/internal/db"
	"github.com/mizuw/laterna/server/cmd/internal/http"
	"github.com/mizuw/laterna/server/cmd/internal/logger"
)

func main() {
	config.AppConfig = config.LoadConfig("config.json")
	logger.InitLoggers()

	go db.ConnectDB()
	http.StartHTTPServer()
}
