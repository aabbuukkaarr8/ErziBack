package cart

type CreateCartDTO struct {
	UserID int `json:"title" validate:"required"`
}

type AddCartItemDTO struct {
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required,min=1"`
}

type UpdateCartItemDTO struct {
	Quantity int `json:"quantity" validate:"required,min=1"`
}

type CartResponseDTO struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}

type CartItemResponseDTO struct {
	ID        int `json:"item_id"`
	CartID    int `json:"cart_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
