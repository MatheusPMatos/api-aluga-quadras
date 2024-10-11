package main

import (
	"fmt"
	"log"

	"github.com/MatheusPMatos/api-aluga-quadras/config"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/infra"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/server"
)

const Version = "0.0.1"

func main() {
	environments := config.GetEnvs()
	db, err := infra.ConectaComBancodeDados(environments)
	if err != nil {
		log.Fatalf("Erro ao conectar com banco de dados, erro: %s", err.Error())
	}
	r := server.NewServer(db, environments)

	r.Run(fmt.Sprintf("%s:%s", environments.ApiHost, environments.ApiPort))
}
