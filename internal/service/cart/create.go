package cart

import "erzi_new/internal/repository/cart"

func (s *Service) CreateCart(userID int) (*cart.Cart, error) {
	return s.repo.Create(userID)
}
