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

type userHandle struct {
	sv        service.User
	validator *validator.Validate
}

// UserInfo implements UserHandler.
func (u *userHandle) UserInfo(c *gin.Context) {
	usr, err := u.sv.GetById(c.GetUint("user"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao buscar user por id")
		return
	}
	c.JSON(http.StatusOK, usr)
}

func (u *userHandle) Delete(c *gin.Context) {
	idStr := c.Params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}

	if err := u.sv.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("erro: %s", err.Error()))
		return
	}

	c.JSON(http.StatusOK, nil)

}

func (u *userHandle) Edit(c *gin.Context) {
	var user types.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid Json")
		return
	}
	err = u.validator.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("validation error: %s", err.Error()))
		return
	}
	usr, err := u.sv.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("erro: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, usr)
}

// GetById implements UserHandler.
func (u *userHandle) GetById(c *gin.Context) {
	idStr := c.Params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}
	usr, err := u.sv.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("erro: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, usr)
}

func (u *userHandle) Create(c *gin.Context) {
	var user types.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid Json")
		return
	}
	err = u.validator.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("validation error: %s", err.Error()))
		return
	}

	usr, err := u.sv.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("erro: %s", err.Error()))
		return
	}
	c.JSON(http.StatusOK, usr)
}

type UserHandler interface {
	Create(c *gin.Context)
	Edit(c *gin.Context)
	GetById(c *gin.Context)
	Delete(c *gin.Context)
	UserInfo(c *gin.Context)
}

func NewUserHandle(serv service.User, validate *validator.Validate) UserHandler {
	return &userHandle{sv: serv, validator: validate}
}
