package utils

import (
	"fmt"
	"time"

	"github.com/MatheusPMatos/api-aluga-quadras/config"
	"github.com/golang-jwt/jwt/v5"
)

type jwtService struct {
	enviroment config.Environments
}

// CreateAccesstoken implements Jwt.
func (j *jwtService) CreateAccesstoken(userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		&jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("%d", userId),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		})

	tokenString, err := token.SignedString([]byte(j.enviroment.TokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// DecodeAccessToken implements Jwt.
func (j *jwtService) DecodeAccessToken(tokens string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.enviroment.TokenSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}
	return nil, fmt.Errorf("invalid token claims")
}

type Jwt interface {
	CreateAccesstoken(userId uint) (string, error)
	DecodeAccessToken(token string) (*jwt.MapClaims, error)
}

func NewJwt(envs config.Environments) Jwt {
	return &jwtService{enviroment: envs}
}
