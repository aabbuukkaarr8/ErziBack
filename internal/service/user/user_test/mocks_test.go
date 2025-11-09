package user

import (
	repo "erzi_new/internal/repository/user"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(u *repo.User) (*repo.User, error) {
	args := m.Called(u)
	return args.Get(0).(*repo.User), args.Error(1)
}
