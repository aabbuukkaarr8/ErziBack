package cart

import (
	repoUser "erzi_new/internal/repository/user"
	product "erzi_new/internal/service/user"
)

type Service interface {
	Create(input product.CreateUser) (*repoUser.User, error)
}
