package product

import (
	"time"

	"erzi_new/internal/repository/product"
)

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

func (m *Product) FillFromDB(dbm *product.Product) {
	m.ID = dbm.ID
	m.Title = dbm.Title
	m.Description = dbm.Description
	m.Price = dbm.Price
	m.ImageURL = dbm.ImageURL
	m.Quantity = dbm.Quantity
	m.Category = dbm.Category
	m.CreatedAt = dbm.CreatedAt
}

func (s *Service) GetByID(id int) (*Product, error) {
	dbp, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	p := &Product{}
	p.FillFromDB(dbp)

	return p, nil
}
