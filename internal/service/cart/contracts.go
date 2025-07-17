package cart

import (
	repoCart "erzi_new/internal/repository/cart"
)

type Repo interface {
	CreateCart(userID int) (*repoCart.Cart, error)
	GetCart(userID int) (int, error)
}
