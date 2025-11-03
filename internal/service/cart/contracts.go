package cart

import (
	"context"
	repoCart "erzi_new/internal/repository/cart"
	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, userID uuid.UUID, status string) (*repoCart.Model, error)
	GetActive(ctx context.Context, userID uuid.UUID) (int, error)
	RestoreCart(ctx context.Context, userID uuid.UUID) error
}
