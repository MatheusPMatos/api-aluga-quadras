package repository

import (
	"time"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"gorm.io/gorm"
)

type reservation struct {
	DB *gorm.DB
}

// GetByDate implements Reservation.
func (r *reservation) GetByDate(date time.Time) ([]types.Reservation, error) {
	panic("Metodo GetByData Não implementado!")
	var reservas = []types.Reservation{}
	err := r.DB.First(&types.Reservation{}).Error
	return reservas, err
}

// Create implements Reservation.
func (r *reservation) Create(reserva types.Reservation) (*types.Reservation, error) {
	err := r.DB.Create(&reserva).Error
	return &reserva, err
}

// GetByProductID implements Reservation.
func (r *reservation) GetByProductID(productId uint) ([]types.Reservation, error) {
	var reservas = []types.Reservation{}
	err := r.DB.Joins("inner join schedules on schedules.product_id = ?", productId).Find(&reservas).Error
	return reservas, err
}

// GetByUserId implements Reservation.
func (r *reservation) GetByUserId(userID uint) ([]types.Reservation, error) {
	var reservar = []types.Reservation{}
	err := r.DB.Where(&types.Reservation{UserID: userID}).Find(&reservar).Error
	return reservar, err
}

// Update implements Reservation.
func (r *reservation) Update(reserva types.Reservation) (*types.Reservation, error) {
	err := r.DB.Model(&types.Reservation{}).Where("id = ?", reserva.ID).
		Updates(map[string]interface{}{
			"is_paid": reserva.IsPaid,
		}).Error
	return &reserva, err
}

// delete implements Reservation.
func (r *reservation) Delete(reservaId uint) error {
	return r.DB.Delete(types.Reservation{}, reservaId).Error
}

type Reservation interface {
	Create(reserva types.Reservation) (*types.Reservation, error)
	Update(reserva types.Reservation) (*types.Reservation, error)
	Delete(reservaId uint) error
	GetByUserId(userID uint) ([]types.Reservation, error)
	GetByProductID(productId uint) ([]types.Reservation, error)
	GetByDate(date time.Time) ([]types.Reservation, error)
}

func NewReservantionRepository(db *gorm.DB) Reservation {
	return &reservation{DB: db}
}
