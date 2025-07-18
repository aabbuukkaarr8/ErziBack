package cartItem

import "erzi_new/internal/repository/cartItem"

func (s *Service) Add(p AddCartItem) (*CartItem, error) {
	existing, err := s.repo.GetByCartAndProduct(p.CartID, p.ProductID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		updated, err := s.repo.UpdateQuantity(existing.ID, existing.Quantity+1)
		if err != nil {
			return nil, err
		}
		return &CartItem{
			ProductID: updated.ProductID,
			CartID:    updated.CartID,
			Quantity:  updated.Quantity,
		}, nil
	}
	toDB := cartItem.CartItem{
		ProductID: p.ProductID,
		CartID:    p.CartID,
	}
	added, err := s.repo.Create(toDB.CartID, toDB.ProductID)
	if err != nil {
		return nil, err
	}
	fromDB := CartItem{
		ProductID: added.ProductID,
		CartID:    added.CartID,
		Quantity:  1,
	}

	return &fromDB, nil
}
