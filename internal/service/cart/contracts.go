package cart

import (
	repoCart "erzi_new/internal/repository/cart"
)

type Repo interface {
	Create(userID int, status string) (*repoCart.Model, error)
	GetActive(userID int) (int, error)
}
