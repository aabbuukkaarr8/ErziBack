package product

import (
	"erzi_new/internal/repository/product"
	"time"
)

func (m *Model) FillFromDB(dbm *product.Model) {
	m.ID = dbm.ID
	m.Title = dbm.Title
	m.Description = dbm.Description
	m.Price = dbm.Price
	m.ImageURL = dbm.ImageURL
	m.IsActive = dbm.IsActive
	m.Category = dbm.Category
	m.CreatedAt = dbm.CreatedAt
	m.BulkDiscountQuantity = dbm.BulkDiscountQuantity
	m.BulkDiscountPrice = dbm.BulkDiscountPrice
}

func (m *Attribute) FillFromDB(dbm *product.Attribute) {
	m.ID = dbm.ID
	m.ProductID = dbm.ProductID
	m.Key = dbm.Key
	m.Value = dbm.Value
}

type AttributeInput struct {
	ProductID int
	Key       string
	Value     string
}

type Attribute struct {
	ID        int
	ProductID int
	Key       string
	Value     string
}

type CreateProduct struct {
	Title                string
	Description          string
	Price                float64
	IsActive             bool
	Category             string
	BulkDiscountQuantity int
	BulkDiscountPrice    float64
}

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

type UpdateProduct struct {
	ID                   int
	Title                *string
	Description          *string
	Price                *float64
	ImageURL             *string
	IsActive             *bool
	Category             *string
	BulkDiscountQuantity *int
	BulkDiscountPrice    *float64
}
