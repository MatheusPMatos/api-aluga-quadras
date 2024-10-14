package infra

import (
	"fmt"
	"log"

	"github.com/MatheusPMatos/api-aluga-quadras/config"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConectaComBancodeDados(envs config.Environments) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		envs.DbHost, envs.DbUser, envs.DbPass, envs.DbName, envs.DbPort)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
		return nil, err
	}

	DB.AutoMigrate(
		&types.User{},
		&types.Product{},
		&types.Schedule{},
		&types.Reservation{},
	)
	return DB, nil

}
