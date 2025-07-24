package user

import (
	"github.com/google/uuid"
	"time"
)

type CreateUser struct {
	Username string
	Email    string
	Password string
	Role     string
}

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}
