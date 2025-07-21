package cart

func (s *Service) GetActive(userID int) (int, error) {
	dbp, err := s.repo.GetActive(userID)
	if err != nil {
		return 0, err
	}
	return dbp, nil
}
