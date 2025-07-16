package cartItem

import (
	"erzi_new/internal/repository/cartItem"
)

type Repo interface {
	Create(cartID, productID int) (*cartItem.CartItem, error)
}
