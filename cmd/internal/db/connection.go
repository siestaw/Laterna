package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/siestaw/laterna/server/cmd/internal/logger"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./db.sql")
	if err != nil {
		logger.DBLogger.Printf("An Error occured while connecting to the database: %v", err)
		return
	}
	logger.DBLogger.Printf("Successfully connected to the database!")
	InitDB()
}

func InitDB() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS lamp_state (
		id TEXT PRIMARY KEY,
		color TEXT NOT NULL,
		updated_at DATETIME NOT NULL
		);
	`)

	if err != nil {
		logger.DBLogger.Fatalf("An error occured while initializing the database: %v", err)
	}
}
