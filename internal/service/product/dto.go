package product

import "time"
import "erzi_new/internal/repository/product"

func (m *Model) FillFromDB(dbm *product.Model) {
	m.ID = dbm.ID
	m.Title = dbm.Title
	m.Description = dbm.Description
	m.Price = dbm.Price
	m.ImageURL = dbm.ImageURL
	m.Quantity = dbm.Quantity
	m.Category = dbm.Category
	m.CreatedAt = dbm.CreatedAt
}

type CreateProduct struct {
	Title       string
	Description string
	Price       float64
	Quantity    int
	Category    string
}

type Model struct {
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
