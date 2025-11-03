package cart

import (
	"context"
	"github.com/google/uuid"
)

func (s *Service) Restore(ctx context.Context, userID uuid.UUID) error {
	err := s.repo.RestoreCart(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
