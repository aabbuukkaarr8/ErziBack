package cartItem

import "errors"

func (s *Service) Decrement(itemID int) (*CartItem, error) {
	// 1. получаем позицию
	existing, err := s.repo.GetByID(itemID)
	if err != nil {
		return nil, err
	}
	if existing.Quantity <= 1 {
		return nil, errors.New("quantity cannot go below 1")
	}
	// 2. обновляем quantity -1
	updated, err := s.repo.UpdateQuantity(itemID, existing.Quantity-1)
	if err != nil {
		return nil, err
	}
	return &CartItem{
		ProductID: updated.ProductID,
		CartID:    updated.CartID,
		Quantity:  updated.Quantity,
	}, nil
}
