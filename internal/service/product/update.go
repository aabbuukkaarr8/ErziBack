package product

import (
	"context"
	repoProduct "erzi_new/internal/repository/product"
)

func (s *Service) Update(ctx context.Context, p UpdateProduct) (*Model, error) {
	current, err := s.repo.GetByID(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	if p.Title != nil {
		current.Title = *p.Title
	}
	if p.Description != nil {
		current.Description = *p.Description
	}
	if p.ImageURL != nil {
		current.ImageURL = *p.ImageURL
	}
	if p.IsActive != nil {
		current.IsActive = *p.IsActive
	}
	if p.Category != nil {
		current.Category = *p.Category
	}
	if p.Prices != nil {
		prices := make(repoProduct.Prices, len(*p.Prices))
		for i, priceEntry := range *p.Prices {
			prices[i] = repoProduct.PriceEntry{
				Quantity: priceEntry.Quantity,
				Price:    priceEntry.Price,
			}
		}
		current.Prices = prices
	}
	updated, err := s.repo.Update(ctx, current)
	if err != nil {
		return nil, err
	}

	fromDb := &Model{}
	fromDb.FillFromDB(updated)
	return fromDb, nil

}
