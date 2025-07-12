package user

import (
	userservice "erzi_new/internal/service/user"
)

type Service interface {
	Create(input userservice.CreateUser) (*userservice.User, error)
	Login(email, password string) (string, error)
}
