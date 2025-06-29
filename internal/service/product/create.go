package product

import (
	"time"

	"erzi_new/internal/repository/product"
)

type CreateProduct struct {
	Title       string
	Description string
	Price       float64
	Quantity    int
}

func (s *Service) Create(p CreateProduct) (*Product, error) {
	toDB := product.Product{
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		CreatedAt:   time.Now(),
	}
	created, err := s.repo.Create(&toDB)
	if err != nil {
		return nil, err
	}

	fromDB := Product{
		ID:          created.ID,
		Title:       created.Title,
		Description: created.Description,
		Price:       created.Price,
		ImageURL:    created.ImageURL,
		Quantity:    created.Quantity,
		CreatedAt:   created.CreatedAt,
	}
	return &fromDB, nil
}
