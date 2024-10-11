package utils

import "github.com/MatheusPMatos/api-aluga-quadras/config"

type jwt struct {
	enviroment config.Environments
}

// CreateAccesstoken implements Jwt.
func (j *jwt) CreateAccesstoken() (string, error) {
	panic("unimplemented")
}

// DecodeAccessToken implements Jwt.
func (j *jwt) DecodeAccessToken(token string) {
	panic("unimplemented")
}

type Jwt interface {
	CreateAccesstoken() (string, error)
	DecodeAccessToken(token string)
}

func NewJwt(envs config.Environments) Jwt {
	return &jwt{enviroment: envs}
}
