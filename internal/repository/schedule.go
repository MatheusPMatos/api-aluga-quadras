package repository

import (
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"gorm.io/gorm"
)

type schedule struct {
	DB *gorm.DB
}

// Update implements Schedule.
func (s *schedule) Update(sched types.Schedule) (*types.Schedule, error) {
	err := s.DB.Model(&types.Schedule{}).Where("id = ?", sched.ID).
		Updates(map[string]interface{}{
			"enable": sched.Enable,
		}).Error
	return &sched, err
}

type Schedule interface {
	Update(sched types.Schedule) (*types.Schedule, error)
}

func NewScheduleRepository(db *gorm.DB) Schedule {
	return &schedule{DB: db}
}
