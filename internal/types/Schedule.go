package types

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	InitialTime  time.Time
	FinalTime    time.Time
	Weekday      time.Weekday
	Enable       bool
	ProductID    uint
	Reservations []Reservation `gorm:"foreignkey:ScheduleID"`
}
