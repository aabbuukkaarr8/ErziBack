package user

import (
	svc "erzi_new/internal/service/user"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Create(p svc.CreateUser) (*svc.User, error) {
	args := m.Called(p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*svc.User), args.Error(1)
}
