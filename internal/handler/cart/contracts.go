package cart

import "context"
import "github.com/google/uuid"

type Service interface {
	GetActive(ctx context.Context, userID uuid.UUID) (int, error)
	Restore(ctx context.Context, userID uuid.UUID) error
}
