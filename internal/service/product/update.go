package product

import "context"

func (s *Service) Update(ctx context.Context, p UpdateProduct) (*Model, error) {
	current, err := s.repo.GetByID(ctx, p.ID)
	if err != nil {
		return nil, err
	}
	if p.Title != nil {
		current.Title = *p.Title
	}
	if p.Description != nil {
		current.Description = *p.Description
	}
	if p.Price != nil {
		current.Price = *p.Price
	}
	if p.ImageURL != nil {
		current.ImageURL = *p.ImageURL
	}
	if p.IsActive != nil {
		current.IsActive = *p.IsActive
	}
	if p.Category != nil {
		current.Category = *p.Category
	}
	updated, err := s.repo.Update(ctx, current)
	if err != nil {
		return nil, err
	}

	fromDb := &Model{}
	fromDb.FillFromDB(updated)
	return fromDb, nil

}
