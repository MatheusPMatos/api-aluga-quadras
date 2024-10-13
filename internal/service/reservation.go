package service

import (
	"errors"
	"fmt"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
)

type reservation struct {
	repo repository.Reservation
}

// Create implements Reservation.
func (r *reservation) Create(reserva types.Reservation) (*types.Reservation, error) {
	reservaJaExistente, err := r.repo.GetByDate(reserva.ID, reserva.Date)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar reserva, erro: %s", err.Error())
	}
	if reservaJaExistente != nil {
		return nil, errors.New("HORARIO RESERVADO")
	}
	return r.repo.Create(reserva)
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
