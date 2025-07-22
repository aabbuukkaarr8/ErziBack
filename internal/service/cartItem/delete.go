package cartItem

func (s *Service) Delete(itemID int) error {
	return s.repo.Delete(itemID)
}
