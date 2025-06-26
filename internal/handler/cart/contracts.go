package cart

import "erzi_new/internal/repository/cart"

type Service interface {
	CreateCart(userID int) (*cart.Cart, error)
}

//type Service interface {
//	CreateCart(userID int) (*cart.Cart, error)
//	GetCart(id int) (*cart.Cart, error)
//	AddToCart(cartID, productID, qty int) (*cart.CartItem, error)
//	UpdateCartItem(itemID, qty int) error
//	RemoveCartItem(itemID int) error
//	ListCartItems(cartID int) ([]cart.CartItem, error)
//}
