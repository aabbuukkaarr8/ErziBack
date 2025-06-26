package product

func (s *Service) GetAll() ([]Product, error) {
	dbProductsPtr, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, err
	}

	products := make([]Product, 0, len(dbProductsPtr))
	for _, dbp := range dbProductsPtr {
		var p Product
		p.FillFromDB(&dbp)
		products = append(products, p)
	}

	return products, nil
}
