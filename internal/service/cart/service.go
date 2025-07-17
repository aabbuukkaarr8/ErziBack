package cart

type Service struct {
	repo Repo
}

func NewService(
	repo Repo,
) *Service {
	return &Service{
		repo: repo,
	}
}
