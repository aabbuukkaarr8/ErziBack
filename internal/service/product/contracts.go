package product

import (
	"context"
	"erzi_new/internal/repository/product"
)

type Repo interface {
	Create(ctx context.Context, p *product.Model) (*product.Model, error)
	GetByID(ctx context.Context, id int) (*product.Model, error)
	GetAllProducts() ([]product.Model, error)
	GetActiveProducts() ([]product.Model, error)
	Update(ctx context.Context, p *product.Model) (*product.Model, error)
	Delete(ctx context.Context, id int) error
	CreateAttributes(ctx context.Context, a *product.Attribute) (*product.Attribute, error)
	GetAttributes(ctx context.Context, id int) ([]product.Attribute, error)
	Hide(ctx context.Context, id int, status bool) error
	GetStatus(ctx context.Context, id int) (bool, error)
}
