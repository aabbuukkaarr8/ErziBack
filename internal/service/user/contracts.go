package user

import (
	"context"
	repoUser "erzi_new/internal/repository/user"
)

type Repo interface {
	Create(ctx context.Context, u *repoUser.User) (*repoUser.User, error)
	GetByEmail(ctx context.Context, email string) (*repoUser.User, error)
}
