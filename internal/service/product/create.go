package product

import (
	"context"
	"time"

	"erzi_new/internal/repository/product"
)

func (s *Service) Create(ctx context.Context, p CreateProduct) (*Model, error) {
	toDB := product.Model{
		Title:                p.Title,
		Description:          p.Description,
		Price:                p.Price,
		Quantity:             p.Quantity,
		Category:             p.Category,
		CreatedAt:            time.Now(),
		BulkDiscountQuantity: p.BulkDiscountQuantity,
		BulkDiscountPrice:    p.BulkDiscountPrice,
	}
	created, err := s.repo.Create(ctx, &toDB)
	if err != nil {
		return nil, err
	}

	fromDB := Model{
		ID:                   created.ID,
		Title:                created.Title,
		Description:          created.Description,
		Price:                created.Price,
		ImageURL:             created.ImageURL,
		Quantity:             created.Quantity,
		Category:             created.Category,
		CreatedAt:            created.CreatedAt,
		BulkDiscountQuantity: created.BulkDiscountQuantity,
		BulkDiscountPrice:    created.BulkDiscountPrice,
	}
	return &fromDB, nil
}
