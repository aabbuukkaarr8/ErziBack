package cartItem

import "context"

func (s *Service) Increment(ctx context.Context, itemID int) (*Model, error) {
	existing, err := s.repo.GetByID(ctx, itemID)
	if err != nil {
		return nil, err
	}
	updated, err := s.repo.UpdateQuantity(ctx, itemID, existing.Quantity+1)
	if err != nil {
		return nil, err
	}
	return &Model{
		ProductID: updated.ProductID,
		CartID:    updated.CartID,
		Quantity:  updated.Quantity,
	}, nil
}
