package types

type ProductType int

const (
	ProductTypeUnknow      = 0
	ProductTypeQuadra      = 1
	ProductTypeEquipamento = 2
)

type Product struct {
	Name        string
	Description string
	ProdType    ProductType
}
