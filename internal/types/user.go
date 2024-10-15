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
	Name         string        `json:"name"`
	Cpf          string        `json:"cpf"`
	Email        string        `json:"email" gorm:"uniqueIndex"`
	Password     string        `json:"password,omitempty"`
	UsrType      UserType      `json:"usr_type"`
	Products     []Product     `json:"-" gorm:"foreignkey:UserID"`
	Reservations []Reservation `json:"-" gorm:"foreignkey:UserID"`
}
