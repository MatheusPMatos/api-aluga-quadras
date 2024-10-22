package types

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	InitialTime  time.Time     `json:"initial_time"`
	FinalTime    time.Time     `json:"final_time"`
	Weekday      time.Weekday  `json:"weekday"`
	Enable       bool          `json:"enable"`
	ProductID    uint          `json:"product_id"`
	Reservations []Reservation `json:"-" gorm:"foreignkey:ScheduleID"`
}
type ScheduleDto struct {
	Reserved bool      `json:"reserved"`
	Date     time.Time `json:"date"`
	Schedule
}
