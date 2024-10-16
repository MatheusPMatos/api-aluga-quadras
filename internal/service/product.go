package service

import (
	"time"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/repository"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
)

type product struct {
	repo repository.Product
}

// Create implements Product.
func (p *product) Create(product types.Product) (*types.Product, error) {
	//SOMENTE VENDEDOR PODE CRIAR
	product.Scheds = createSchedule()
	return p.repo.Create(product)
}

func createSchedule() []types.Schedule {
	var scheds []types.Schedule
	for i := 0; i < 7; i++ {
		for j := 0; j < 24; j++ {
			scheds = append(scheds, types.Schedule{
				InitialTime: time.Date(1899, 12, 30, j, 0, 0, 0, time.UTC),
				FinalTime:   time.Date(1899, 12, 30, j, 59, 0, 0, time.UTC),
				Weekday:     time.Weekday(i),
				Enable:      j > 6,
			})
		}
	}
	return scheds
}

// Delete implements Product.
func (p *product) Delete(productId uint) error {
	//SOMENTE VENDEDOR PODE DELETAR
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
	//SOMENTE VENDEDOR PODE ALTERAR
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
