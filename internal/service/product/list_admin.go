package product

func (s *Service) GetAll() ([]Model, error) {
	dbProductsPtr, err := s.repo.GetAllProducts()
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
