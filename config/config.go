package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Environments struct {
	ApiPort string `env:"API_PORT"`
	ApiHost string `env:"API_HOST"`

	TokenSecret string `env:"TOKEN_SECRET"`

	DbHost string `env:"DB_HOST"`
	DbPort string `env:"DB_PORT"`
	DbUser string `env:"DB_USER"`
	DbPass string `env:"DB_PASS"`
	DbName string `env:"DB_NAME"`
}

func GetEnvs() Environments {
	var configs Environments
	err := env.Parse(&configs)
	if err != nil {
		log.Fatalf("Error to Set Environments")
	}
	return configs
}
