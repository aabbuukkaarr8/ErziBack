package cartItem

import "erzi_new/internal/repository/cartItem"

type AddCartItem struct {
	ProductID int
	CartID    int
}

type CartItem struct {
	ProductID int
	CartID    int
	Quantity  int
}

func (s *Service) Add(p AddCartItem) (*CartItem, error) {
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
