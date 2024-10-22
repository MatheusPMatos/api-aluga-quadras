package repository

import (
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"gorm.io/gorm"
)

type product struct {
	DB *gorm.DB
}

// Create implements Product.
func (p *product) Create(product types.Product) (*types.Product, error) {
	err := p.DB.Create(&product).Error
	return &product, err
}

// Delete implements Product.
func (p *product) Delete(productId uint) error {
	return p.DB.Delete(&types.Product{}, productId).Error
}

// GetAll implements Product.
func (p *product) GetAll() ([]types.Product, error) {
	var products = []types.Product{}
	err := p.DB.Find(&products).Error
	return products, err
}

func (p *product) GetById(productId uint) (*types.Product, error) {
	var product types.Product
	err := p.DB.First(&product, productId).Error
	return &product, err
}

// Update implements Product.
func (p *product) Update(product types.Product) (*types.Product, error) {
	err := p.DB.Model(&types.Product{}).Where("id = ?", product.ID).
		Updates(map[string]interface{}{
			"name":        product.Name,
			"description": product.Description,
			"prod_type":   product.ProdType,
			"price":       product.Price,
		}).Error
	return &product, err
}

type Product interface {
	Create(product types.Product) (*types.Product, error)
	GetById(productId uint) (*types.Product, error)
	GetAll() ([]types.Product, error)
	Delete(productId uint) error
	Update(product types.Product) (*types.Product, error)
}

func NewProductRepository(db *gorm.DB) Product {
	return &product{DB: db}
}
