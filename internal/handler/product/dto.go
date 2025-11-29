package product

import (
	"time"
)

// PriceEntry представляет одну запись цены
type PriceEntry struct {
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type CreateProduct struct {
	Title       string       `json:"title" validate:"required"`
	Description string       `json:"description" validate:"required"`
	IsActive    bool         `json:"is_active"`
	Category    string       `json:"category" validate:"required,oneof=honey-jam meltwater mineral-water equipment"`
	Prices      []PriceEntry `json:"prices" validate:"required,min=1,dive"`
}

type Model struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	ImageURL    string       `json:"image_url"`
	IsActive    bool         `json:"is_active"`
	Category    string       `json:"category"`
	CreatedAt   time.Time    `json:"created_at"`
	Prices      []PriceEntry `json:"prices"`
	Attribute   []Attribute  `json:"params"`
}

type UpdateProduct struct {
	Title       *string       `json:"title"`
	Description *string       `json:"description"`
	ImageURL    *string       `json:"image_url"`
	IsActive    *bool         `json:"is_active"`
	Category    *string       `json:"category"`
	Prices      *[]PriceEntry `json:"prices"`
}
type CreateAttributeRequest struct {
	ProductID int    `json:"product_id"`
	Key       string `json:"key" binding:"required"`
	Value     string `json:"value" binding:"required"`
}

type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
