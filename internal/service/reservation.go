package service

import (
	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
)

type reservation struct {
	repo repository.Reservation
}

// Create implements Reservation.
func (r *reservation) Create(reserva types.Reservation) (*types.Reservation, error) {
	panic("unimplemented")
}

// GetByProductID implements Reservation.
func (r *reservation) GetByProductID(productId uint) ([]types.Reservation, error) {
	return r.repo.GetByProductID(productId)
}

// GetByUserId implements Reservation.
func (r *reservation) GetByUserId(userID uint) ([]types.Reservation, error) {
	return r.repo.GetByUserId(userID)
}

// Update implements Reservation.
func (r *reservation) Update(reserva types.Reservation) (*types.Reservation, error) {
	return r.repo.Update(reserva)
}

// delete implements Reservation.
func (r *reservation) Delete(reservaId uint) error {
	return r.repo.Delete(reservaId)
}

type Reservation interface {
	Create(reserva types.Reservation) (*types.Reservation, error)
	Update(reserva types.Reservation) (*types.Reservation, error)
	Delete(reservaId uint) error
	GetByUserId(userID uint) ([]types.Reservation, error)
	GetByProductID(productId uint) ([]types.Reservation, error)
}

func NewReservantionService(repo repository.Reservation) Reservation {
	return &reservation{repo: repo}
}
