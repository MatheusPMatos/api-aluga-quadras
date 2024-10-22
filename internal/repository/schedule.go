package repository

import (
	"time"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"gorm.io/gorm"
)

type schedule struct {
	DB *gorm.DB
}

// GetByProductIdAndDate implements Schedule.
func (s *schedule) GetByProductIdAndDate(productId uint, weekday time.Weekday) ([]types.Schedule, error) {
	var schudules = []types.Schedule{}
	err := s.DB.Where(&types.Schedule{Weekday: weekday, ProductID: productId}).Find(&schudules).Error
	return schudules, err

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
	GetByProductIdAndDate(productId uint, weekday time.Weekday) ([]types.Schedule, error)
}

func NewScheduleRepository(db *gorm.DB) Schedule {
	return &schedule{DB: db}
}
