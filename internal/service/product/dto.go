package product

import (
	"erzi_new/internal/repository/product"
	"time"
)

// PriceEntry представляет одну запись цены
type PriceEntry struct {
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (m *Model) FillFromDB(dbm *product.Model) {
	m.ID = dbm.ID
	m.Title = dbm.Title
	m.Description = dbm.Description
	m.ImageURL = dbm.ImageURL
	m.IsActive = dbm.IsActive
	m.Category = dbm.Category
	m.CreatedAt = dbm.CreatedAt
	m.Prices = make([]PriceEntry, len(dbm.Prices))
	for i, p := range dbm.Prices {
		m.Prices[i] = PriceEntry{
			Quantity: p.Quantity,
			Price:    p.Price,
		}
	}
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
	Title       string
	Description string
	IsActive    bool
	Category    string
	Prices      []PriceEntry
}

type Model struct {
	ID          int
	Title       string
	Description string
	ImageURL    string
	IsActive    bool
	Category    string
	CreatedAt   time.Time
	Prices      []PriceEntry
}

type UpdateProduct struct {
	ID          int
	Title       *string
	Description *string
	ImageURL    *string
	IsActive    *bool
	Category    *string
	Prices      *[]PriceEntry
}
