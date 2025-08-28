package cart

import "github.com/google/uuid"

func (s *Service) Restore(userID uuid.UUID) error {
	err := s.repo.RestoreCart(userID)
	if err != nil {
		return err
	}
	return nil
}
