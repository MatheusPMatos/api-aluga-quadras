package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Environments struct {
	ApiPort string `env:"API_PORT"`
}

func GetEnvs() Environments {
	var configs Environments
	err := env.Parse(&configs)
	if err != nil {
		log.Fatalf("Error to Set Environments")
	}
	return configs
}
