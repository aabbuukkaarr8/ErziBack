package cart

import (
	repoCart "erzi_new/internal/repository/cart"
	"github.com/google/uuid"
)

type Repo interface {
	Create(userID uuid.UUID, status string) (*repoCart.Model, error)
	GetActive(userID uuid.UUID) (int, error)
}
