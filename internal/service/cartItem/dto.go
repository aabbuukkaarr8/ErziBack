package cartItem

import (
	"erzi_new/internal/repository/cartItem"
	"github.com/google/uuid"
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
	m.Product.BulkDiscountQuantity = dbm.Product.BulkDiscountQuantity
	m.Product.BulkDiscountPrice = dbm.Product.BulkDiscountPrice
}

type AddCartItemRequest struct {
	UserID    uuid.UUID
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
	Title                string  `json:"title"`
	Price                float64 `json:"price"`
	ImageURL             string  `json:"image_url"`
	BulkDiscountQuantity int     `json:"bulk_discount_quantity"`
	BulkDiscountPrice    float64 `json:"bulk_discount_price"`
}

type ModelResponse struct {
	ID        int             `json:"id"`
	CartID    int             `json:"cart_id"`
	ProductID int             `json:"product_id"`
	Quantity  int             `json:"quantity"`
	CreatedAt time.Time       `json:"created_at"`
	Product   ProductMiniInfo `json:"product"`
}
