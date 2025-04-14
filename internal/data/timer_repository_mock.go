package data

import (
	"database/sql"
)

type TimerRepositoryMock struct {
	DB *sql.DB
}

func NewTimerRepositoryMock(db *sql.DB) *TimerRepositoryMock {
	return &TimerRepositoryMock{DB: db}
}

// IMPLEMENTATIONS

func (r *TimerRepositoryMock) GetAllTimers() ([]Timer, error) {
	var timers []Timer

	timers = append(timers, Timer{
		ID:            1,
		Name:          "Pomodoro",
		Description:   "25 minutes of focus followed by 5 minutes of rest",
		FocusDuration: 25,
		RestDuration:  5,
	})

	timers = append(timers, Timer{
		ID:            2,
		Name:          "Pomodoro small",
		Description:   "15 minutes of focus followed by 15 minutes of rest",
		FocusDuration: 15,
		RestDuration:  15,
	})

	return timers, nil
}
