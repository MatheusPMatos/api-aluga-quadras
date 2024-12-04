package service

import (
	"errors"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/utils"
)

type user struct {
	repo repository.User
}

// Delete implements User.
func (u *user) Delete(userId uint) error {
	return u.repo.Delete(userId)
}

// Create implements User.
func (u *user) Create(user types.User) (*types.User, error) {
	user.Password = utils.ShaEncode(user.Password)
	if user.UsrType != 1 && user.UsrType != 2 {
		return nil, errors.New("invalid user profile")
	}
	return u.repo.Create(user)
}

// GetById implements User.
func (u *user) GetById(userId uint) (*types.User, error) {
	return u.repo.GetById(userId)
}

// Update implements User.
func (u *user) Update(user types.User) (*types.User, error) {
	return u.repo.Update(user)
}

type User interface {
	Create(user types.User) (*types.User, error)
	GetById(userId uint) (*types.User, error)
	Update(user types.User) (*types.User, error)
	Delete(userId uint) error
}

func NewUserService(rp repository.User) User {
	return &user{repo: rp}

}
