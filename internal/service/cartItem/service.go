package cartItem

type Service struct {
	repo        Repo
	cartService CartService
}

func NewService(
	repo Repo,
	cartService CartService,
) *Service {
	return &Service{
		repo:        repo,
		cartService: cartService,
	}
}
