package cartItem

import (
	"erzi_new/internal/service/cartItem"
)

type Service interface {
	Add(p cartItem.AddCartItemRequest) (*cartItem.Model, error)
	GetAll(cartID int) ([]cartItem.ModelResponse, error)
	Increment(ItemID int) (*cartItem.Model, error)
	Decrement(ItemID int) (*cartItem.Model, error)
	Delete(itemID int) error
	DeleteAll(cartID int) error
}

type CartService interface {
	GetActive(userID int) (int, error)
}
