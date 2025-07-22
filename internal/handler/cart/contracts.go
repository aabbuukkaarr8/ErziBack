package cart

import "github.com/google/uuid"

type Service interface {
	GetActive(userID uuid.UUID) (int, error)
}
