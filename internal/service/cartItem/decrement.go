package cartItem

import (
	"context"
	"errors"
)

func (s *Service) Decrement(ctx context.Context, itemID int) (*Model, error) {
	existing, err := s.repo.GetByID(ctx, itemID)
	if err != nil {
		return nil, err
	}
	if existing.Quantity <= 0 {
		err := s.Delete(ctx, itemID)
		if err != nil {
			return nil, err
		}
	}
	if existing.Quantity == 1 {
		return nil, errors.New("quantity cannot go below 1")
	}
	updated, err := s.repo.UpdateQuantity(ctx, itemID, existing.Quantity-1)
	if err != nil {
		return nil, err
	}
	return &Model{
		ProductID: updated.ProductID,
		CartID:    updated.CartID,
		Quantity:  updated.Quantity,
	}, nil
}
