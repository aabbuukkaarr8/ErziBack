package product

import "time"

type Product struct {
	ProductID   int
	Title       string
	Description string
	Price       float64
	ImageURL    string
	Quantity    int
	CreatedAt   time.Time
}
