package cartItem

import (
	"erzi_new/internal/service/cartItem"
)

type Service interface {
	Add(p cartItem.AddCartItem) (*cartItem.CartItem, error)
	GetAll(cartID int) ([]cartItem.CartItemResponse, error)
}
