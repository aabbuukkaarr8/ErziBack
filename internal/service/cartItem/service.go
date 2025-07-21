package cartItem

import cartService "erzi_new/internal/service/cart"

type Service struct {
	repo     Repo
	cartRepo cartService.Repo
}

func NewService(
	repo Repo,
	cartRepo cartService.Repo,
) *Service {
	return &Service{
		repo:     repo,
		cartRepo: cartRepo,
	}
}
