package data

type TimerRepository interface {
	//GetTimer(id int64) (Timer, error)
	GetAllTimers() ([]Timer, error)
	// CreateTimer(timer Timer) (int64, error)
	// UpdateTimer(timer Timer) error
	// DeleteTImer(id int64) error
}
