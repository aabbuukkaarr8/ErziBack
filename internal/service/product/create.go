package product

import (
	"time"

	"erzi_new/internal/repository/product"
)

func (s *Service) Create(p CreateProduct) (*Product, error) {
	toDB := product.Product{
		Title:       p.Title,
		Description: p.Description,
		Price:       p.Price,
		Quantity:    p.Quantity,
		Category:    p.Category,
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
		Category:    created.Category,
		CreatedAt:   created.CreatedAt,
	}
	return &fromDB, nil
}
