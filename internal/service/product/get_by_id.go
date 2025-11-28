package product

import "context"

func (s *Service) GetByID(ctx context.Context, id int) (*Model, error) {
	dbp, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	p := &Model{}
	p.FillFromDB(dbp)

	return p, nil
}
