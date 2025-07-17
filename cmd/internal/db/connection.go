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
		DeleteController(0)
	}
}

func InitDB() {
	_, err := DB.Exec(` 
		CREATE TABLE IF NOT EXISTS controllers (
		id TEXT PRIMARY KEY,
		token_hash TEXT,
		color TEXT,
		updated_at DATETIME
		);
	`)

	if err != nil {
		logger.DBLogger.Fatalf("An error occured while initializing the database: %v", err)
	}

}