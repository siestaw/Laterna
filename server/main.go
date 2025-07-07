package main

import (
	"github.com/mizuw/laterna/server/database"
	"github.com/mizuw/laterna/server/globals"
	"github.com/mizuw/laterna/server/http"
)

func main() {
	globals.AppConfig = globals.LoadConfig("config.jsonc")
	globals.InitLoggers()

	go database.ConnectDB()
	http.StartHTTPServer()
}
