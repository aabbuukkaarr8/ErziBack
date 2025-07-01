package product

import (
	svc "erzi_new/internal/service/product"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Create(p svc.CreateProduct) (*svc.Product, error) {
	args := m.Called(p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*svc.Product), args.Error(1)
}

func (m *MockService) GetAll() ([]svc.Product, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]svc.Product), args.Error(1)
}

func (m *MockService) GetByID(id int) (*svc.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*svc.Product), args.Error(1)
}

func (m *MockService) Update(p svc.UpdateProduct) (*svc.Product, error) {
	args := m.Called(p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*svc.Product), args.Error(1)
}

func (m *MockService) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
