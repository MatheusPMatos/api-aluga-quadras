package main

import (
	"fmt"

	"github.com/MatheusPMatos/api-aluga-quadras/config"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/server"
)

const Version = "0.0.1"

func main() {
	environments := config.GetEnvs()
	println(environments.ApiPort)

	r := server.NewServer()

	r.Run(fmt.Sprintf("%s:%s", environments.ApiHost, environments.ApiPort))
}
