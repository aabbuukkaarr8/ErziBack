package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}
