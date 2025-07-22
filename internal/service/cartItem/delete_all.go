package cartItem

func (s *Service) DeleteAll(userID int) error {
	activeCart, err := s.cartService.GetActive(userID)
	if err != nil {
		return err
	}

	return s.repo.DeleteAll(activeCart)
}
