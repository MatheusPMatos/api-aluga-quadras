package handlers

import (
	"net/http"
	"strconv"

	"github.com/MatheusPMatos/api-aluga-quadras/internal/service"
	"github.com/MatheusPMatos/api-aluga-quadras/internal/types"
	"github.com/gin-gonic/gin"
)

type product struct {
	sv service.Product
}

// Update implements Product.
func (p *product) Update(c *gin.Context) {
	var product types.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid Json")
		return
	}

	usr, err := p.sv.Update(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao criar produto")
		return
	}
	c.JSON(http.StatusOK, usr)
}

// Create implements Product.
func (p *product) Create(c *gin.Context) {
	var product types.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid Json")
		return
	}

	usr, err := p.sv.Create(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao criar produto")
		return
	}
	c.JSON(http.StatusOK, usr)
}

// Delete implements Product.
func (p *product) Delete(c *gin.Context) {
	idStr := c.Params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}
	if err := p.sv.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao alterar produto")
		return
	}
	c.JSON(http.StatusOK, nil)
}

// Edit implements Product.
func (p *product) Edit(c *gin.Context) {

}

// GetAll implements Product.
func (p *product) GetAll(c *gin.Context) {
	usr, err := p.sv.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao buscar produtos")
		return
	}
	c.JSON(http.StatusOK, usr)
}

// GetById implements Product.
func (p *product) GetById(c *gin.Context) {
	idStr := c.Params.ByName("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid params")
		return
	}
	usr, err := p.sv.GetById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "erro ao alterar produto")
		return
	}
	c.JSON(http.StatusOK, usr)
}

type Product interface {
	Create(c *gin.Context)
	Edit(c *gin.Context)
	GetById(c *gin.Context)
	Delete(c *gin.Context)
	GetAll(c *gin.Context)
	Update(c *gin.Context)
}

func NewProductHandler(serv service.Product) Product {
	return &product{sv: serv}
}
