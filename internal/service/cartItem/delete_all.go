package cartItem

import (
	"context"
	"github.com/google/uuid"
)

func (s *Service) DeleteAll(ctx context.Context, userID uuid.UUID) error {
	activeCart, err := s.cartService.GetActive(ctx, userID)
	if err != nil {
		return err
	}

	return s.repo.DeleteAll(ctx, activeCart)
}
