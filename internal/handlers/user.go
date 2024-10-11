package handlers

import (
	"net/http"
	"strconv"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/service"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"github.com/gin-gonic/gin"
)

type userHandle struct {
	sv service.User
}

func (u *userHandle) Delete(c *gin.Context) {
	idStr := c.Params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}

	if err := u.sv.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao deletar user")
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

	usr, err := u.sv.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao alterar user")
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
		c.JSON(http.StatusInternalServerError, "erro ao alterar user")
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

	usr, err := u.sv.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao alterar user")
		return
	}
	c.JSON(http.StatusOK, usr)
}

type UserHandler interface {
	Create(c *gin.Context)
	Edit(c *gin.Context)
	GetById(c *gin.Context)
	Delete(c *gin.Context)
}

func NewUserHandle(serv service.User) UserHandler {
	return &userHandle{sv: serv}
}
