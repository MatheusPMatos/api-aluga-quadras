package repository

import "github.com/MatheusPMatos/api-aluga-quadras/internal/types"

type user struct {
	//DB *gorm.DB
}

// Create implements User.
func (u *user) Create(user types.User) (types.User, error) {
	panic("unimplemented")
}

// Delete implements User.
func (u *user) Delete(userID uint) error {
	panic("unimplemented")
}

// GetById implements User.
func (u *user) GetById(userId uint) (types.User, error) {
	panic("unimplemented")
}

// Update implements User.
func (u *user) Update(user types.UserType) (types.User, error) {
	panic("unimplemented")
}

type User interface {
	Create(user types.User) (types.User, error)
	Update(user types.UserType) (types.User, error)
	Delete(userID uint) error
	GetById(userId uint) (types.User, error)
}

func NewUserRepository() User {
	return &user{}
}
