package cartItem

import (
	"erzi_new/internal/repository/cartItem"
)

type Repo interface {
	Create(cartID, productID int) (*cartItem.CartItem, error)
	GetAll(cartId int) ([]cartItem.CartItemWithProduct, error)
	UpdateQuantity(id, quantity int) (*cartItem.CartItem, error)
	GetByCartAndProduct(cartID, productID int) (*cartItem.CartItem, error)
	GetByID(id int) (*cartItem.CartItem, error)
	Delete(id int) error
}
