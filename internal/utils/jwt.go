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

type userClains struct {
	Sub     uint
	SubName string
	Claims  jwt.Claims
}

// CreateAccesstoken implements Jwt.
func (j *jwtService) CreateAccesstoken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": "name",
			"user_id":  "name",
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(j.enviroment)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// DecodeAccessToken implements Jwt.
func (j *jwtService) DecodeAccessToken(tokens string) error {
	token, err := jwt.Parse(tokens, func(token *jwt.Token) (interface{}, error) {
		return j.enviroment.ApiHost, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

type Jwt interface {
	CreateAccesstoken() (string, error)
	DecodeAccessToken(token string) error
}

func NewJwt(envs config.Environments) Jwt {
	return &jwtService{enviroment: envs}
}
