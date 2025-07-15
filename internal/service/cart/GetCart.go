package cart

func (s *Service) GetCart(userID int) (int, error) {
	dbp, err := s.repo.GetCart(userID)
	if err != nil {
		return 0, err
	}
	return dbp, nil
}
