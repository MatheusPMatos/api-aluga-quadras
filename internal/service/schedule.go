package service

import (
	"time"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
)

type schedule struct {
	repo    repository.Schedule
	reserRP repository.Reservation
}

// GetByProductWeekDay implements Schedule.
func (s *schedule) GetByProductWeekDay(produdctId uint, date time.Time) ([]types.ScheduleDto, error) {
	var schedsDto = []types.ScheduleDto{}

	scheds, err := s.repo.GetByProductIdAndDate(produdctId, date.Weekday())
	if err != nil {
		return nil, err
	}

	reservas, err := s.reserRP.GetByProductAndDate(produdctId, date)
	if err != nil {
		return nil, err
	}

	for _, sched := range scheds {
		for _, reserva := range reservas {
			if sched.ID == reserva.ScheduleID {
				schedsDto = append(schedsDto, types.ScheduleDto{
					Date:     date,
					Reserved: true,
					Schedule: sched,
				})
				goto next
			}
		}
		schedsDto = append(schedsDto, types.ScheduleDto{
			Date:     date,
			Reserved: false,
			Schedule: sched,
		})
	next:
	}
	return schedsDto, nil
}

type Schedule interface {
	GetByProductWeekDay(produdctId uint, date time.Time) ([]types.ScheduleDto, error)
}

func NewScheduleService(rp repository.Schedule, reserRP repository.Reservation) Schedule {
	return &schedule{repo: rp, reserRP: reserRP}
}
