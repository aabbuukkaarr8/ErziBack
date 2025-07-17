package cartItem

func (s *Service) Increment(itemID int) (*CartItem, error) {
	existing, err := s.repo.GetByID(itemID)
	if err != nil {
		return nil, err
	}
	updated, err := s.repo.UpdateQuantity(itemID, existing.Quantity+1)
	if err != nil {
		return nil, err
	}
	return &CartItem{
		ProductID: updated.ProductID,
		CartID:    updated.CartID,
		Quantity:  updated.Quantity,
	}, nil
}
