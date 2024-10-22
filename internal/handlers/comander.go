package handlers

import (
	"github.com/MatheusPMatos/api-aluga-quadras/config"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/service"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/utils"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Comander struct {
	User       UserHandler
	Products   Product
	Midlewares service.Midleware
	Auth       Auth
	Rservation Reservation
}

func NewComander(db *gorm.DB, envs config.Environments) Comander {
	userRepo := repository.NewUserRepository(db)
	validate := validator.New(validator.WithRequiredStructEnabled())

	return Comander{
		User:       NewUserHandle(service.NewUserService(userRepo), validate),
		Products:   NewProductHandler(service.NewProductService(repository.NewProductRepository(db), userRepo), validate),
		Midlewares: service.NewMidleware(userRepo, utils.NewJwt(envs)),
		Auth: NewAuthHandle(
			service.NewAuthService(utils.NewJwt(envs), userRepo),
		),
		Rservation: NewReservationHandle(service.NewReservantionService(repository.NewReservantionRepository(db)), validate),
	}

}
