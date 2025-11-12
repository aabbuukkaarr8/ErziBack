package product

import "time"

type Model struct {
	ID                   int
	Title                string
	Description          string
	Price                float64
	ImageURL             string
	IsActive             bool
	Category             string
	CreatedAt            time.Time
	BulkDiscountQuantity int
	BulkDiscountPrice    float64
}

type Attribute struct {
	ID        int
	ProductID int
	Key       string
	Value     string
}
