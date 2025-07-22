package cartItem

import "github.com/google/uuid"

func (s *Service) DeleteAll(userID uuid.UUID) error {
	activeCart, err := s.cartService.GetActive(userID)
	if err != nil {
		return err
	}

	return s.repo.DeleteAll(activeCart)
}
