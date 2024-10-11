package repository

import (
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"gorm.io/gorm"
)

type user struct {
	DB *gorm.DB
}

// Create implements User.
func (u *user) Create(user types.User) (*types.User, error) {
	result := u.DB.Create(&user)
	return &user, result.Error
}

// Delete implements User.
func (u *user) Delete(userID uint) error {
	result := u.DB.Delete(&types.User{}, userID)
	return result.Error
}

// GetById implements User.
func (u *user) GetById(userId uint) (*types.User, error) {
	var user types.User
	err := u.DB.First(&user, userId).Error
	return &user, err
}

// Update implements User.
func (u *user) Update(user types.User) (*types.User, error) {

	err := u.DB.Model(&types.User{}).Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"name":     user.Name,
			"cpf":      user.Cpf,
			"email":    user.Email,
			"usr_type": user.UsrType,
			"password": user.Password}).Error
	return &user, err
}

type User interface {
	Create(user types.User) (*types.User, error)
	Update(user types.User) (*types.User, error)
	Delete(userID uint) error
	GetById(userId uint) (*types.User, error)
}

func NewUserRepository(db *gorm.DB) User {
	return &user{DB: db}
}
