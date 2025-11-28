package product

func (s *Service) List() ([]Model, error) {
	dbProductsPtr, err := s.repo.GetActiveProducts()
	if err != nil {
		return nil, err
	}

	products := make([]Model, 0, len(dbProductsPtr))
	for _, dbp := range dbProductsPtr {
		var p Model
		p.FillFromDB(&dbp)
		products = append(products, p)
	}

	return products, nil
}
