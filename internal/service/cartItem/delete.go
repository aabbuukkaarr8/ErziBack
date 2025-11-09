package cartItem

import "context"

func (s *Service) Delete(ctx context.Context,
	itemID int) error {
	return s.repo.Delete(ctx, itemID)
}
