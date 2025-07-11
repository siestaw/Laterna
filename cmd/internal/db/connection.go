package db

import (
	"database/sql"
	"flag"

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

	resetAdmin := flag.Bool("resetAdminToken", false, "recreate the admin token")
	flag.Parse()

	if *resetAdmin {
		ResetAdmin()
	}
}

func InitDB() {
	// Legacy purposes, DELETE
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

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS controllers (
		id INTEGER PRIMARY KEY,
		token_hash TEXT NOT NULL,
		color TEXT,
		updated_at DATETIME
	);`)
	if err != nil {
		logger.DBLogger.Fatalf("An error occured while initializing the database: %v", err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS permissions (
		controller_id INTEGER,
		target_id INTEGER
	);`)
	if err != nil {
		logger.DBLogger.Fatalf("An error occured while initializing the database: %v", err)
	}
}
