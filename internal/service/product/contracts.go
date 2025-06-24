package product

import "erzi_new/internal/repository/product"

type Repo interface {
	Create(p *product.Product) (*product.Product, error)
	GetByID(id int) (*product.Product, error)
	GetAllProducts() ([]product.Product, error)
}
