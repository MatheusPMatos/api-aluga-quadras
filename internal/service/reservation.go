package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"gorm.io/gorm"
)

type reservation struct {
	repo repository.Reservation
}

// Create implements Reservation.
func (r *reservation) Create(reserva types.Reservation) (*types.Reservation, error) {
	reservaJaExistente, err := r.repo.GetByDate(reserva.ScheduleID, reserva.Date)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			goto procede
		}
		return nil, fmt.Errorf("erro ao gerar reserva, erro: %s", err.Error())
	}
procede:
	if reservaJaExistente.ID != 0 && reservaJaExistente.ScheduleID == reserva.ScheduleID {
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
	reservaAtual, err := r.repo.GetById(reserva.ID)
	if err != nil {
		return nil, err
	}
	if err := validaAlteracaoExclusao(*reservaAtual); err != nil {
		return nil, err
	}
	return r.repo.Update(reserva)
}

// delete implements Reservation.
func (r *reservation) Delete(reservaId uint) error {
	reserva, err := r.repo.GetById(reservaId)
	if err != nil {
		return err
	}
	if err := validaAlteracaoExclusao(*reserva); err != nil {
		return err
	}

	return r.repo.Delete(reservaId)
}

func validaAlteracaoExclusao(reserva types.Reservation) error {
	if reserva.IsPaid {
		return errors.New("reserva paga, nao e possivel excluir")
	}
	if time.Until(reserva.Date) < (time.Hour * 24) {
		return errors.New("nao e possivel excluir a reserva a menos de 24h")
	}
	return nil
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
