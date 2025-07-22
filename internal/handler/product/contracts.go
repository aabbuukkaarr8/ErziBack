package product

import "erzi_new/internal/service/product"

type Service interface {
	Create(p product.CreateProduct) (*product.Model, error)
	GetByID(id int) (*product.Model, error)
	GetAll() ([]product.Model, error)
	Update(p product.UpdateProduct) (*product.Model, error)
	Delete(id int) error
}
