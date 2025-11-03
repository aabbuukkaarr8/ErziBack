package user

import (
	"context"
	userservice "erzi_new/internal/service/user"
)

type Service interface {
	Create(ctx context.Context, input userservice.CreateUser) (*userservice.User, error)
	Login(ctx context.Context, email, password string) (string, error)
}
