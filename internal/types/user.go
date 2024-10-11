package types

import "gorm.io/gorm"

type UserType int

const (
	UserTypeUnknown   = 0
	UserTypeVendedor  = 1
	UserTypeComprador = 2
)

type User struct {
	gorm.Model
	Name     string
	Cpf      string
	Email    string
	Password string
	UsrType  UserType
	Products []Product `gorm:"foreignkey:UserID"`
}
