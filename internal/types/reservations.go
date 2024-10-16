package types

import (
	"time"

	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	ScheduleID uint      `json:"schedule_id" validate:"required,number"`
	Date       time.Time `json:"date" validate:"required,datetime"`
	UserID     uint      `json:"user_id" validate:"required,number"`
	IsPaid     bool      `json:"is_paid"`
}
