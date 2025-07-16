package product

import (
	"errors"
	"testing"
	"time"

	repo "erzi_new/internal/repository/product"
	svc "erzi_new/internal/service/product"

	"github.com/stretchr/testify/assert"
)

func TestService_GetByID_Success(t *testing.T) {
	mr := new(MockRepo)
	service := svc.NewService(mr)

	now := time.Now().Truncate(time.Second)
	repoProd := &repo.Product{
		ID:          7,
		Title:       "Tasty Water",
		Description: "Pure spring",
		Price:       2.50,
		ImageURL:    "http://img",
		Quantity:    100,
		Category:    "beverages",
		CreatedAt:   now,
	}

	mr.On("GetByID", 7).Return(repoProd, nil)

	got, err := service.GetByID(7)
	assert.NoError(t, err)

	want := &svc.Product{}
	want.FillFromDB(repoProd)
	assert.Equal(t, want, got)

	mr.AssertExpectations(t)
}

func TestService_GetByID_Error(t *testing.T) {
	mr := new(MockRepo)
	service := svc.NewService(mr)

	mr.On("GetByID", 99).Return((*repo.Product)(nil), errors.New("not found"))

	got, err := service.GetByID(99)
	assert.Nil(t, got)
	assert.EqualError(t, err, "not found")

	mr.AssertExpectations(t)
}
