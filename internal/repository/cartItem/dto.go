package cartItem

import "time"

type CartItem struct {
	ID        int
	CartID    int
	ProductID int
	Quantity  int
	CreatedAt time.Time
}

type CartItemWithProduct struct {
	ID        int
	CartID    int
	ProductID int
	Quantity  int
	CreatedAt time.Time
	Product   struct {
		Title    string
		Price    float64
		ImageURL string
	}
}
