package cartItem

import (
	"erzi_new/internal/repository/cartItem"
	"time"
)

func (m *ModelResponse) FillFromDB(dbm *cartItem.ModelWithProduct) {
	m.ID = dbm.ID
	m.Quantity = dbm.Quantity
	m.CartID = dbm.CartID
	m.ProductID = dbm.ProductID
	m.CreatedAt = dbm.CreatedAt
	m.Product.Title = dbm.Product.Title
	m.Product.Price = dbm.Product.Price
	m.Product.ImageURL = dbm.Product.ImageURL
}

type AddCartItemRequest struct {
	UserID    int
	ProductID int
	CartID    int
}

type Model struct {
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

type ModelResponse struct {
	ID        int             `json:"id"`
	CartID    int             `json:"cart_id"`
	ProductID int             `json:"product_id"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"created_at"`
	Product   ProductMiniInfo `json:"product"`
}
