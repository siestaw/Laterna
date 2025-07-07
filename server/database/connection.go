package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mizuw/laterna/server/globals"
)

func ConnectDB() {
	db, err := sql.Open("sqlite3", "./db.sql")
	if err != nil {
		globals.DBLogger.Printf("An Error occured while connecting to the database: %v", err)
	}
	globals.DBLogger.Printf("Connected")
	defer db.Close()
}
