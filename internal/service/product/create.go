package product

import (
	"context"
	"time"

	"erzi_new/internal/repository/product"
)

func (s *Service) Create(ctx context.Context, p CreateProduct) (*Model, error) {
	prices := make(product.Prices, len(p.Prices))
	for i, priceEntry := range p.Prices {
		prices[i] = product.PriceEntry{
			Quantity: priceEntry.Quantity,
			Price:    priceEntry.Price,
		}
	}
	toDB := product.Model{
		Title:       p.Title,
		Description: p.Description,
		IsActive:    p.IsActive,
		Category:    p.Category,
		CreatedAt:   time.Now(),
		Prices:      prices,
	}
	created, err := s.repo.Create(ctx, &toDB)
	if err != nil {
		return nil, err
	}

	fromDB := Model{}
	fromDB.FillFromDB(created)
	return &fromDB, nil
}
