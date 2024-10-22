package repository

import (
	"time"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"gorm.io/gorm"
)

type reservation struct {
	DB *gorm.DB
}

// GetByProductAndDate implements Reservation.
func (r *reservation) GetByProductAndDate(productId uint, date time.Time) ([]types.Reservation, error) {
	var reservas = []types.Reservation{}
	err := r.DB.Joins("inner join schedules on schedules.id = reservations.schedule_id and schedules.productId = ?", productId).
		Where("DATE(reservation.date) = DATE(?)", date).Find(&reservas).Error
	return reservas, err
}

// GetById implements Reservation.
func (r *reservation) GetById(id uint) (*types.Reservation, error) {
	var reserva types.Reservation
	err := r.DB.First(&reserva, id).Error
	return &reserva, err
}

// GetByDate implements Reservation.
func (r *reservation) GetByDate(scheduleID uint, date time.Time) (*types.Reservation, error) {
	var reservas = types.Reservation{}
	err := r.DB.
		Where(&types.Reservation{ScheduleID: scheduleID}).
		Where("DATE(created_at) = DATE(?)", date).First(&reservas).Error
	return &reservas, err
}

// Create implements Reservation.
func (r *reservation) Create(reserva types.Reservation) (*types.Reservation, error) {
	err := r.DB.Create(&reserva).Error
	return &reserva, err
}

// GetByProductID implements Reservation.
func (r *reservation) GetByProductID(productId uint) ([]types.Reservation, error) {
	var reservas = []types.Reservation{}
	err := r.DB.Joins("inner join schedules on schedules.id = reservations.id and schedules.product_id = ?", productId).Group("reservations.id").Find(&reservas).Error
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
	return r.DB.Delete(&types.Reservation{}, reservaId).Error
}

type Reservation interface {
	Create(reserva types.Reservation) (*types.Reservation, error)
	Update(reserva types.Reservation) (*types.Reservation, error)
	Delete(reservaId uint) error
	GetByUserId(userID uint) ([]types.Reservation, error)
	GetById(id uint) (*types.Reservation, error)
	GetByProductID(productId uint) ([]types.Reservation, error)
	GetByDate(scheduleID uint, date time.Time) (*types.Reservation, error)
	GetByProductAndDate(productId uint, date time.Time) ([]types.Reservation, error)
}

func NewReservantionRepository(db *gorm.DB) Reservation {
	return &reservation{DB: db}
}
