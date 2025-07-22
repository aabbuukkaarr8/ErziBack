package cartItem

import (
	repoCart "erzi_new/internal/repository/cart"
	"erzi_new/internal/repository/cartItem"
)

type Repo interface {
	Create(cartID, productID int) (*cartItem.Model, error)
	GetAll(cartId int) ([]cartItem.ModelWithProduct, error)
	UpdateQuantity(id, quantity int) (*cartItem.Model, error)
	GetByCartAndProduct(cartID, productID int) (*cartItem.Model, error)
	GetByID(id int) (*cartItem.Model, error)
	Delete(id int) error
	DeleteAll(cartID int) error
}

type CartService interface {
	Create(userID int, status string) (*repoCart.Model, error)
	GetActive(userID int) (int, error)
}
