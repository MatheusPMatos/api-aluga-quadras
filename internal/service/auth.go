package service

import (
	"errors"
	"strings"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/dto"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/utils"
	"gorm.io/gorm"
)

type auth struct {
	jwt   utils.Jwt
	usrRp repository.User
}

// Login implements Auth.
func (a *auth) Login(auth dto.Auth) (*dto.Tokens, error) {
	loginErr := errors.New("email ou senha invalido")
	usr, err := a.usrRp.GetByEmail(auth.Email)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, loginErr
		default:
			return nil, err
		}
	}
	if strings.Compare(usr.Password, utils.ShaEncode(auth.Password)) != 0 {
		return nil, loginErr
	}
	token, err := a.jwt.CreateAccesstoken(usr.ID)
	if err != nil {
		return nil, err
	}
	return &dto.Tokens{AccessToken: token}, nil
}

type Auth interface {
	Login(auth dto.Auth) (*dto.Tokens, error)
}

func NewAuthService(jw utils.Jwt, userrepo repository.User) Auth {
	return &auth{jwt: jw, usrRp: userrepo}
}
