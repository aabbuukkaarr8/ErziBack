package cartItem

import "time"

type AddCartItem struct {
	UserID    int
	ProductID int
	CartID    int
}

type CartItem struct {
	ID        int
	ProductID int
	CartID    int
	Quantity  int
	CreatedAt time.Time
}

type ProductMiniInfo struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}

type CartItemResponse struct {
	ID        int             `json:"id"`
	CartID    int             `json:"cart_id"`
	ProductID int             `json:"product_id"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"created_at"`
	Product   ProductMiniInfo `json:"product"`
}
