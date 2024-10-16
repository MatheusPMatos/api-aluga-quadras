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
	Name         string        `json:"name" validate:"required"`
	Cpf          string        `json:"cpf" validate:"required,max=15"`
	Email        string        `json:"email" gorm:"uniqueIndex" validate:"required,email"`
	Password     string        `json:"password,omitempty" validate:"required,min=6"`
	UsrType      UserType      `json:"usr_type" validate:"gte=1,lte=2"`
	Products     []Product     `json:"-" gorm:"foreignkey:UserID"`
	Reservations []Reservation `json:"-" gorm:"foreignkey:UserID"`
}
