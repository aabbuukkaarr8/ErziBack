package cartItem

import (
	"erzi_new/internal/service/cartItem"
)

type Service interface {
	Add(p cartItem.AddCartItem) (*cartItem.CartItem, error)
	GetAll(cartID int) ([]cartItem.CartItemResponse, error)
	Increment(ItemID int) (*cartItem.CartItem, error)
	Decrement(ItemID int) (*cartItem.CartItem, error)
	Delete(itemID int) error
	DeleteAll(cartID int) error
}
