package cartItem

import (
	"erzi_new/internal/service/cartItem"
	"github.com/google/uuid"
)

type Service interface {
	Add(p cartItem.AddCartItemRequest) (*cartItem.Model, error)
	GetAll(cartID int) ([]cartItem.ModelResponse, error)
	Increment(ItemID int) (*cartItem.Model, error)
	Decrement(ItemID int) (*cartItem.Model, error)
	Delete(itemID int) error
	DeleteAll(userID uuid.UUID) error
}

type CartService interface {
	GetActive(userID uuid.UUID) (int, error)
}
