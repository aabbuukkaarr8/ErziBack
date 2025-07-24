package cart

import "github.com/google/uuid"

type Model struct {
	ID     int
	UserID uuid.UUID
	Status string
}
