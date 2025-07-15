package product

import (
	"errors"
	"testing"
	"time"

	repo "erzi_new/internal/repository/product"
	svc "erzi_new/internal/service/product"

	"github.com/stretchr/testify/assert"
)

func TestService_GetAll_Success(t *testing.T) {
	mr := new(MockRepo)
	service := svc.NewService(mr)
	now := time.Now()
	repoProds := []repo.Product{
		{ID: 1, Title: "A", Description: "Desc A", Price: 1.1, ImageURL: "u1", Quantity: 10, Category: "cat1", CreatedAt: now},
		{ID: 2, Title: "B", Description: "Desc B", Price: 2.2, ImageURL: "u2", Quantity: 20, Category: "cat2", CreatedAt: now},
	}
	mr.On("GetAllProducts").Return(repoProds, nil)
	got, err := service.GetAll()
	assert.NoError(t, err)
	want := []svc.Product{
		{ID: 1, Title: "A", Description: "Desc A", Price: 1.1, ImageURL: "u1", Quantity: 10, Category: "cat1", CreatedAt: now},
		{ID: 2, Title: "B", Description: "Desc B", Price: 2.2, ImageURL: "u2", Quantity: 20, Category: "cat2", CreatedAt: now},
	}

	assert.Equal(t, want, got)
	mr.AssertExpectations(t)
}

func TestService_GetAll_Fail(t *testing.T) {
	mr := new(MockRepo)
	service := svc.NewService(mr)
	mr.On("GetAllProducts").Return(nil, errors.New("db error"))
	got, err := service.GetAll()
	assert.Nil(t, got)
	assert.EqualError(t, err, "db error")

	mr.AssertExpectations(t)
}
