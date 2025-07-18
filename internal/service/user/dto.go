package user

import "time"

type CreateUser struct {
	Username string
	Email    string
	Password string
	Role     string
}

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}
