package cartItem

type AddCartItem struct {
	ProductID int
	CartID    int
}

type CartItem struct {
	ProductID int
	CartID    int
	Quantity  int
}
