package product

import "erzi_new/internal/service/product"

type Service interface {
	Create(p product.CreateProduct) (*product.Product, error)
	GetByID(id int) (*product.Product, error)
}
