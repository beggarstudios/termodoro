package data

import (
	"database/sql"
	"time"
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

func (s Store) Init() error {
	var err error

	s.conn, err = sql.Open("sqlite3", "./profiles.db")
	if err != nil {
		return err
	}

	createTableStmt := `
		CREATE TABLE IF NOT EXISTS profiles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL
	);`

	if _, err = s.conn.Exec(createTableStmt); err != nil {
		return err
	}

	return nil
}

func (s *Store) getProfiles() ([]Profile, error) {
	rows, err := s.conn.Query("SELECT * FROM profiles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	profiles := []Profile{}
	for rows.Next() {
		var p Profile
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}

func (s *Store) saveProfile(profile Profile) error {
	if profile.ID == 0 {
		profile.ID = time.Now().UnixNano()
	}

	query := `INSERT INTO profiles (id, title) VALUES (?, ?)`

	if _, err := s.conn.Exec(query, profile.ID, profile.Name); err != nil {
		return err
	}

	return nil
}
