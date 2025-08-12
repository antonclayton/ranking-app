package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes and returns a connection to the SQLite database.
func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CreateTables creates the necessary database tables if they don't already exist.
func CreateTables(db *sql.DB) error {
	placeTable := `
	CREATE TABLE IF NOT EXISTS places (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		types TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	productTable := `
	CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT,
		place_id INTEGER,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(place_id) REFERENCES places(id)
	);
	`

	ratingTable := `
	CREATE TABLE IF NOT EXISTS ratings (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		target_id INTEGER NOT NULL,
		target_type TEXT NOT NULL, -- 'place' or 'product'
		score INTEGER NOT NULL,
		comment TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(placeTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(productTable)
	if err != nil {
		return err
	}

	_, err = db.Exec(ratingTable)
	if err != nil {
		return err
	}

	return nil
}
