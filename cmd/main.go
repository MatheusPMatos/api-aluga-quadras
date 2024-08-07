package main

import "github.com/MatheusPMatos/api-aluga-quadras/config"

const Version = "0.0.1"

func main() {
	environments := config.GetEnvs()
	println(environments.ApiPort)
}
