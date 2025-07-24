package cart

import "github.com/google/uuid"

func (s *Service) GetActive(userID uuid.UUID) (int, error) {
	dbp, err := s.repo.GetActive(userID)
	if err != nil {
		return 0, err
	}
	return dbp, nil
}
