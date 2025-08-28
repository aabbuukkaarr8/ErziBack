package product

import (
	"errors"
	"erzi_new/internal/repository/product"
)

func (s *Service) CreateAttributes(a AttributeInput) (*Attribute, error) {
	toDB := product.Attribute{
		ProductID: a.ProductID,
		Key:       a.Key,
		Value:     a.Value,
	}
	if a.ProductID <= 0 {
		return nil, errors.New("product_id must be > 0")
	}
	if len(a.Key) == 0 {
		return nil, errors.New("key is required")
	}
	if len(a.Value) == 0 {
		return nil, errors.New("value is required")
	}

	_, err := s.repo.GetByID(a.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	created, err := s.repo.CreateAttributes(&toDB)
	if err != nil {
		return nil, err
	}
	fromDB := Attribute{
		ID:        created.ID,
		ProductID: created.ProductID,
		Key:       created.Key,
		Value:     created.Value,
	}
	return &fromDB, nil
}
