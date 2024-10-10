package types

import "time"

type Schedule struct {
	InitialTime time.Time
	FinalTime   time.Time
	Weekday     time.Weekday
	Enable      bool
	ProductID   uint
}
