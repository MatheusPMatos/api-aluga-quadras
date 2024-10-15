package handlers

import (
	"net/http"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/dto"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/service"
	"github.com/gin-gonic/gin"
)

type auth struct {
	sv service.Auth
}

// Login implements Auth.
func (a *auth) Login(c *gin.Context) {
	var login dto.Auth

	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid Json")
		return
	}

	tokens, err := a.sv.Login(login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, tokens)
}

type Auth interface {
	Login(c *gin.Context)
}

func NewAuthHandle(serv service.Auth) Auth {
	return &auth{sv: serv}
}
