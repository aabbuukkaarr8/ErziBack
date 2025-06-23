package product

func (s *Service) GetAll() ([]Product, error) {
	dbProductsPtr, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	dbProducts := dbProductsPtr
	products := make([]Product, 0, len(dbProducts))
	for _, dbp := range dbProducts {
		var p Product
		p.FillFromDB(&dbp)
		products = append(products, p)
	}

	return products, nil
}
