package product

type UpdateProduct struct {
	ID          int
	Title       *string
	Description *string
	Price       *float64
	ImageURL    *string
	Quantity    *int
	Category    *string
}

func (s *Service) Update(p UpdateProduct) (*Product, error) {
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

	fromDb := &Product{}
	fromDb.FillFromDB(updated)
	return fromDb, nil

}
