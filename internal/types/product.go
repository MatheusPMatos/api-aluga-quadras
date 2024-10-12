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
	Name        string
	Description string
	ProdType    ProductType
	UserID      uint
	Scheds      []Schedule `gorm:"foreignkey:ProductID"`
}
