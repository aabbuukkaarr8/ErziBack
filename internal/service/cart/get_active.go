package cart

import (
	"context"
	"github.com/google/uuid"
)

func (s *Service) GetActive(ctx context.Context, userID uuid.UUID) (int, error) {
	dbp, err := s.repo.GetActive(ctx, userID)
	if err != nil {
		return 0, err
	}
	return dbp, nil
}
