package data

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Profile struct {
	ID       int64
	Name     string
	Settings []Setting
}

type Setting struct {
	ID   int64
	Name string
}

type Store struct {
	conn *sql.DB
}

func InitializeStore(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	// TODO revert to store?
	//store := &Store{conn: db}
	if err := EnsureTablesExist(db); err != nil {
		return nil, err
	}

	return db, nil
}

func EnsureTablesExist(db *sql.DB) error {
	// Profiles table

	if err := EnsureProfileTableExists(db); err != nil {
		return err
	}

	return nil
}

func EnsureProfileTableExists(db *sql.DB) error {
	_, err := db.Exec(
		`CREATE TABLE IF NOT EXISTS profiles 
		(id INTEGER PRIMARY KEY, name TEXT NOT NULL)`,
	)
	if err != nil {
		return err
	}

	return nil
}

// func (s *Store) getProfiles() ([]Profile, error) {
// 	rows, err := s.conn.Query("SELECT * FROM profiles")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	profiles := []Profile{}
// 	for rows.Next() {
// 		var p Profile
// 		if err := rows.Scan(&p.ID, &p.Name); err != nil {
// 			return nil, err
// 		}
// 		profiles = append(profiles, p)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return profiles, nil
// }

// func (s *Store) saveProfile(profile Profile) error {
// 	if profile.ID == 0 {
// 		profile.ID = time.Now().UnixNano()
// 	}

// 	query := `INSERT INTO profiles (id, title) VALUES (?, ?)`

// 	if _, err := s.conn.Exec(query, profile.ID, profile.Name); err != nil {
// 		return err
// 	}

// 	return nil
// }
