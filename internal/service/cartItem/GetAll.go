package cartItem

func (s *Service) GetAll(cartID int) ([]CartItemResponse, error) {
	dbitems, err := s.repo.GetAll(cartID)
	if err != nil {
		return nil, err
	}
	cartItems := make([]CartItemResponse, 0, len(dbitems))
	for _, dbitem := range dbitems {
		var cartItem CartItemResponse
		cartItem.FillFromDB(&dbitem)
		cartItems = append(cartItems, cartItem)

	}

	return cartItems, nil
}
