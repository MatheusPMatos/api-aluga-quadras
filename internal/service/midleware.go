package service

import (
	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/utils"
	"github.com/gin-gonic/gin"
)

type midleware struct {
	usrRepo repository.User
	jwt     utils.Jwt
}

// Auth implements Midleware.
func (m *midleware) Auth() gin.HandlerFunc {
	panic("unimplemented")
}

// AuthAdm implements Midleware.
func (m *midleware) AuthAdm() gin.HandlerFunc {
	panic("unimplemented")
}

type Midleware interface {
	Auth() gin.HandlerFunc
	AuthAdm() gin.HandlerFunc
}

func NewMidleware(usrRp repository.User, jwt utils.Jwt) Midleware {
	return &midleware{usrRepo: usrRp, jwt: jwt}
}
