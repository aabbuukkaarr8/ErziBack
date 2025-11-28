package product

import "context"

func (s *Service) Hide(ctx context.Context, id int) error {
	status, err := s.repo.GetStatus(ctx, id)
	if err != nil {
		return err
	}
	err = s.repo.Hide(ctx, id, status)
	if err != nil {
		return err
	}
	return nil
}
