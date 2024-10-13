package types

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	ScheduleID uint
	Date       time.Time
	UserID     uint
	IsPaid     bool
}
