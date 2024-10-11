package service

import (
	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
)

type product struct {
	repo repository.Product
}

// Create implements Product.
func (p *product) Create(product types.Product) (*types.Product, error) {
	return p.repo.Create(product)
}

// Delete implements Product.
func (p *product) Delete(productId uint) error {
	return p.repo.Delete(productId)
}

// GetAll implements Product.
func (p *product) GetAll() ([]types.Product, error) {
	return p.repo.GetAll()
}

// GetById implements Product.
func (p *product) GetById(productId uint) (*types.Product, error) {
	return p.repo.GetById(productId)
}

// Update implements Product.
func (p *product) Update(product types.Product) (*types.Product, error) {
	return p.repo.Update(product)
}

type Product interface {
	Create(product types.Product) (*types.Product, error)
	GetById(productId uint) (*types.Product, error)
	GetAll() ([]types.Product, error)
	Delete(productId uint) error
	Update(product types.Product) (*types.Product, error)
}

func NewProductService(rp repository.Product) Product {
	return &product{repo: rp}
}
