package product

import "context"

func (s *Service) GetAttributes(ctx context.Context, id int) ([]Attribute, error) {
	dbp, err := s.repo.GetAttributes(ctx, id)
	if err != nil {
		return nil, err
	}
	attributes := make([]Attribute, 0, len(dbp))
	for _, dba := range dbp {
		var attribute Attribute
		attribute.FillFromDB(&dba)
		attributes = append(attributes, attribute)
	}
	return attributes, nil

}
