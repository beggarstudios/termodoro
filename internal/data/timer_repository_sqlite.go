package data

import (
	"database/sql"
)

type TimerRepositorySQLite struct {
	DB *sql.DB
}

func NewTimerRepositorySQLite(db *sql.DB) *TimerRepositorySQLite {
	return &TimerRepositorySQLite{DB: db}
}

// IMPLEMENTATIONS

func (r *TimerRepositorySQLite) GetAllTimers() ([]Timer, error) {
	rows, err := r.DB.Query("SELECT id, name, description, focus_duration, rest_duration FROM timers")
	if err != nil {
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	defer rows.Close()

	var timers []Timer
	for rows.Next() {
		var t Timer
		if err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.FocusDuration, &t.RestDuration); err != nil {
			return nil, err
		}
		timers = append(timers, t)
	}

	return timers, nil
}
