package product

import "erzi_new/internal/service/product"
import "context"

type Service interface {
	Create(ctx context.Context, p product.CreateProduct) (*product.Model, error)
	GetByID(ctx context.Context, id int) (*product.Model, error)
	GetAll() ([]product.Model, error)
	Update(ctx context.Context, p product.UpdateProduct) (*product.Model, error)
	Delete(ctx context.Context, id int) error
	CreateAttributes(ctx context.Context, a product.AttributeInput) (*product.Attribute, error)
	GetAttributes(ctx context.Context, id int) ([]product.Attribute, error)
	Hide(ctx context.Context, id int) error
}
