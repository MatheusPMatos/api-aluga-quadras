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

type MyClaims struct {
	Sub                  uint      `json:"sub"` // User ID
	ExpiresAt            time.Time `json:"exp"` // Tempo de expiração
	jwt.RegisteredClaims           // Embeds the standard registered claims
}

// CreateAccesstoken implements Jwt.
func (j *jwtService) CreateAccesstoken(userId uint) (string, error) {

	claim := MyClaims{
		Sub:       userId,
		ExpiresAt: time.Now().Add(time.Hour * 1),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(j.enviroment.TokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// DecodeAccessToken implements Jwt.
func (j *jwtService) DecodeAccessToken(tokens string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokens, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.enviroment.TokenSecret), nil
	})
	if err != nil {
		return 0, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims.Sub, nil
	}
	return 0, fmt.Errorf("invalid token claims")
}

type Jwt interface {
	CreateAccesstoken(userId uint) (string, error)
	DecodeAccessToken(token string) (uint, error)
}

func NewJwt(envs config.Environments) Jwt {
	return &jwtService{enviroment: envs}
}
