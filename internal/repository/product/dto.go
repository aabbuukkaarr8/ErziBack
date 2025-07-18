package product

import "time"

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
