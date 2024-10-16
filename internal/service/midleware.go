package service

import (
	"net/http"
	"strconv"
	"strings"

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
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		token := strings.TrimPrefix(bearerToken, "Bearer ")

		claim, err := m.jwt.DecodeAccessToken(token)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		id, err := claim.GetSubject()
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		idd, err := strconv.Atoi(id)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.Set("user", idd)
	}
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
