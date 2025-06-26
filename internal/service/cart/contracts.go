package cart

import "erzi_new/internal/repository/cart"

type Repo interface {
	Create(userID int) (*cart.Cart, error)
}
