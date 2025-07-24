// internal/service/product/mocks_test.go
package product

import (
	repo "erzi_new/internal/repository/product"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Create(p *repo.Model) (*repo.Model, error) {
	args := m.Called(p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Model), args.Error(1)
}
func (m *MockRepo) GetAllProducts() ([]repo.Model, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]repo.Model), args.Error(1)
}
func (m *MockRepo) GetByID(id int) (*repo.Model, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Model), args.Error(1)
}
func (m *MockRepo) Update(p *repo.Model) (*repo.Model, error) {
	args := m.Called(p)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repo.Model), args.Error(1)
}
func (m *MockRepo) Delete(id int) error { return m.Called(id).Error(0) }
