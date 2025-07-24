package product

func (s *Service) Update(p UpdateProduct) (*Model, error) {
	current, err := s.repo.GetByID(p.ID)
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
	if p.Quantity != nil {
		current.Quantity = *p.Quantity
	}
	if p.Category != nil {
		current.Category = *p.Category
	}
	updated, err := s.repo.Update(current)
	if err != nil {
		return nil, err
	}

	fromDb := &Model{}
	fromDb.FillFromDB(updated)
	return fromDb, nil

}
