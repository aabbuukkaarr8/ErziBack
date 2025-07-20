package product

func (s *Service) GetByID(id int) (*Product, error) {
	dbp, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	p := &Product{}
	p.FillFromDB(dbp)

	return p, nil
}
