package types

import "gorm.io/gorm"

type ProductType int

const (
	ProductTypeUnknow      = 0
	ProductTypeQuadra      = 1
	ProductTypeEquipamento = 2
)

type Product struct {
	gorm.Model
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"max=240"`
	ProdType    ProductType `json:"prod_type" validate:"gte=1,lte=2"`
	UserID      uint        `json:"user_id"`
	Price       float64     `json:"price" validate:"required,numeric"`
	Scheds      []Schedule  `json:"-" gorm:"foreignkey:ProductID"`
}
