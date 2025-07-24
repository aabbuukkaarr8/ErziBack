package product

import "erzi_new/internal/repository/product"

type Repo interface {
	Create(p *product.Model) (*product.Model, error)
	GetByID(id int) (*product.Model, error)
	GetAllProducts() ([]product.Model, error)
	Update(p *product.Model) (*product.Model, error)
	Delete(id int) error
}
