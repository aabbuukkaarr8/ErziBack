package cartItem

func (s *Service) DeleteAll(cartID int) error {
	return s.repo.DeleteAll(cartID)
}
