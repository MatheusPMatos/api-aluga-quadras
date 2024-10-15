package handlers

import (
	"github.com/MatheusPMatos/api-aluga-quadras/config"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/service"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/utils"
	"gorm.io/gorm"
)

type Comander struct {
	User       UserHandler
	Products   Product
	Midlewares service.Midleware
	Auth       Auth
}

func NewComander(db *gorm.DB, envs config.Environments) Comander {
	userRepo := repository.NewUserRepository(db)

	return Comander{
		User:       NewUserHandle(service.NewUserService(userRepo)),
		Products:   NewProductHandler(service.NewProductService(repository.NewProductRepository(db))),
		Midlewares: service.NewMidleware(userRepo, utils.NewJwt(envs)),
		Auth: NewAuthHandle(
			service.NewAuthService(utils.NewJwt(envs), userRepo),
		),
	}

}
