package product

func (s *Service) GetByID(id int) (*Model, error) {
	dbp, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	p := &Model{}
	p.FillFromDB(dbp)

	return p, nil
}
