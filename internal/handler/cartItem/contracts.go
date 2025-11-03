package cartItem

import (
	"context"
	"erzi_new/internal/service/cartItem"
	"github.com/google/uuid"
)

type Service interface {
	Add(ctx context.Context, p cartItem.AddCartItemRequest) (*cartItem.Model, error)
	GetAll(ctx context.Context, cartID int) ([]cartItem.ModelResponse, error)
	Increment(ctx context.Context, ItemID int) (*cartItem.Model, error)
	Decrement(ctx context.Context, ItemID int) (*cartItem.Model, error)
	Delete(ctx context.Context, itemID int) error
	DeleteAll(ctx context.Context, userID uuid.UUID) error
}

type CartService interface {
	GetActive(ctx context.Context, userID uuid.UUID) (int, error)
}
