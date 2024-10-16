package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/service"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type reservation struct {
	sv       service.Reservation
	validate *validator.Validate
}

// Delete implements Reservation.
func (r *reservation) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}
	if err := r.sv.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao deletar produto")
		return
	}
	c.JSON(http.StatusOK, nil)
}

// Create implements Reservation.
func (r *reservation) Create(c *gin.Context) {
	var reserva types.Reservation

	err := c.ShouldBindJSON(&reserva)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid Json")
		return
	}
	err = r.validate.Struct(reserva)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("validation error: %s", err.Error()))
		return
	}

	usr, err := r.sv.Create(reserva)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao criar produto")
		return
	}
	c.JSON(http.StatusOK, usr)
}

// Edit implements Reservation.
func (r *reservation) Edit(c *gin.Context) {
	var reserva types.Reservation

	err := c.ShouldBindJSON(&reserva)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid Json")
		return
	}
	err = r.validate.Struct(reserva)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("validation error: %s", err.Error()))
		return
	}

	usr, err := r.sv.Update(reserva)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao criar produto")
		return
	}
	c.JSON(http.StatusOK, usr)
}

// GetByProduct implements Reservation.
func (r *reservation) GetByProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}
	reservas, err := r.sv.GetByProductID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao deletar produto")
		return
	}
	c.JSON(http.StatusOK, reservas)
}

// GetByUser implements Reservation.
func (r *reservation) GetByUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}
	reservas, err := r.sv.GetByUserId(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao deletar produto")
		return
	}
	c.JSON(http.StatusOK, reservas)
}

type Reservation interface {
	Create(c *gin.Context)
	Edit(c *gin.Context)
	GetByProduct(c *gin.Context)
	GetByUser(c *gin.Context)
	Delete(c *gin.Context)
}

func NewReservationHandle(serv service.Reservation, validator *validator.Validate) Reservation {
	return &reservation{sv: serv, validate: validator}
}
