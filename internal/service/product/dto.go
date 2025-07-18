package product

import "time"

type CreateProduct struct {
	Title       string
	Description string
	Price       float64
	Quantity    int
	Category    string
}

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImageURL    string
	Quantity    int
	Category    string
	CreatedAt   time.Time
}

type UpdateProduct struct {
	ID          int
	Title       *string
	Description *string
	Price       *float64
	ImageURL    *string
	Quantity    *int
	Category    *string
}
