package cartItem

import (
	"context"
	repoCart "erzi_new/internal/repository/cart"
	"erzi_new/internal/repository/cartItem"
	"github.com/google/uuid"
)

type Repo interface {
	Create(ctx context.Context, cartID, productID int) (*cartItem.Model, error)
	GetAll(ctx context.Context, cartId int) ([]cartItem.ModelWithProduct, error)
	UpdateQuantity(ctx context.Context, id, quantity int) (*cartItem.Model, error)
	GetByCartAndProduct(ctx context.Context, cartID, productID int) (*cartItem.Model, error)
	GetByID(ctx context.Context, id int) (*cartItem.Model, error)
	Delete(ctx context.Context, id int) error
	DeleteAll(ctx context.Context, cartID int) error
}

type CartService interface {
	Create(ctx context.Context, userID uuid.UUID, status string) (*repoCart.Model, error)
	GetActive(ctx context.Context, userID uuid.UUID) (int, error)
}
