package product

import "time"

type CreateProduct struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=1"`
	Quantity    int     `json:"quantity" validate:"min=0"`
	Category    string  `json:"category" validate:"required,oneof=honey-jam meltwater mineral-water equipment"`
}

type Product struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	ImageURL    string    `json:"image_url"`
	Quantity    int       `json:"quantity"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}

type UpdateProduct struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Quantity    *int     `json:"quantity"`
	ImageURL    *string  `json:"image_url"`
	Category    *string  `json:"category"`
}
