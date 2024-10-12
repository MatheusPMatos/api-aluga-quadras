package types

import "gorm.io/gorm"

type Reservation struct {
	gorm.Model
	ScheduleID uint
	UserID     uint
	IsPaid     bool
}
