package cartItem

import "context"

func (s *Service) GetAll(ctx context.Context, cartID int) ([]ModelResponse, error) {
	dbItems, err := s.repo.GetAll(ctx, cartID)
	if err != nil {
		return nil, err
	}
	cartItems := make([]ModelResponse, 0, len(dbItems))
	for _, dbItem := range dbItems {
		var cartItem ModelResponse
		cartItem.FillFromDB(&dbItem)
		cartItems = append(cartItems, cartItem)

	}

	return cartItems, nil
}
