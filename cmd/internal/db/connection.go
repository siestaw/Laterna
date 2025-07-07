package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/siestaw/laterna/server/cmd/internal/logger"
)

func ConnectDB() {
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		logger.DBLogger.Printf("An Error occured while connecting to the database: %v", err)
	}
	logger.DBLogger.Printf("Connected")
	defer db.Close()
}
