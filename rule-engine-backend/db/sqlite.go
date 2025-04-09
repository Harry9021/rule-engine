package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() (*sql.DB, error) {
	// log.Println("🔌 Connecting to SQLite database...")
	db, err := sql.Open("sqlite3", "./db/rule.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	// Create rules table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS rules (
		id TEXT PRIMARY KEY,
		condition TEXT NOT NULL,
		action TEXT NOT NULL
	);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create rules table: %v", err)
	}

	// fmt.Println("✅ Connected to SQLite database!")
	return db, nil
}
