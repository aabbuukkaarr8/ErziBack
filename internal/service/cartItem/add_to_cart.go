package cartItem

import (
	"context"
	"erzi_new/internal/repository/cartItem"
)

func (s *Service) Add(ctx context.Context, p AddCartItemRequest) (*Model, error) {
	cartID, err := s.cartService.GetActive(ctx, p.UserID)
	if err != nil {
		cart, err := s.cartService.Create(ctx, p.UserID, "active")
		if err != nil {
			return nil, err
		}
		cartID = cart.ID

	}
	p.CartID = cartID

	existing, err := s.repo.GetByCartAndProduct(ctx, p.CartID, p.ProductID)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		updated, err := s.repo.UpdateQuantity(ctx, existing.ID, existing.Quantity+1)
		if err != nil {
			return nil, err
		}
		return &Model{
			ProductID: updated.ProductID,
			CartID:    updated.CartID,
			Quantity:  updated.Quantity,
		}, nil
	}
	toDB := cartItem.Model{
		ProductID: p.ProductID,
		CartID:    p.CartID,
	}
	added, err := s.repo.Create(ctx, toDB.CartID, toDB.ProductID)
	if err != nil {
		return nil, err
	}
	fromDB := Model{
		ProductID: added.ProductID,
		CartID:    added.CartID,
		Quantity:  1,
	}

	return &fromDB, nil
}
