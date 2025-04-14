package data

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func OpenDatabaseConnection(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := EnsureTablesExist(db); err != nil {
		return nil, err
	}

	return db, nil
}

func EnsureTablesExist(db *sql.DB) error {

	if err := EnsureTimerTableExists(db); err != nil {
		return err
	}

	return nil
}

func EnsureTimerTableExists(db *sql.DB) error {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS timers (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			focus_duration INTEGER,
			rest_duration INTEGER
		);`,
	)

	if err != nil {
		return err
	}

	return nil
}
