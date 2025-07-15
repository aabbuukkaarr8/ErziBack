package user

import (
	repoUser "erzi_new/internal/repository/user"
)

type Repo interface {
	Create(u *repoUser.User) (*repoUser.User, error)
	GetByEmail(email string) (*repoUser.User, error)
}
