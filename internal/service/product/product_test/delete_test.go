package product

import (
	"errors"
	"testing"
	"time"

	repo "erzi_new/internal/repository/product"
	svc "erzi_new/internal/service/product"

	"github.com/stretchr/testify/assert"
)

func TestService_Delete_Succes(t *testing.T) {
	mr := new(MockRepo)
	service := svc.NewService(mr)
	now := time.Now()
	existing := &repo.Model{
		ID:          7,
		Title:       "A",
		Description: "B",
		Price:       1.23,
		ImageURL:    "u",
		Quantity:    1,
		Category:    "beverages",
		CreatedAt:   now,
	}
	mr.On("GetByID", 7).Return(existing, nil)
	mr.On("Delete", 7).Return(nil)

	err := service.Delete(7)
	assert.NoError(t, err)

	mr.AssertExpectations(t)
}

func TestService_Delete_GetByID_NotFound(t *testing.T) {
	mr := new(MockRepo)
	service := svc.NewService(mr)

	mr.On("GetByID", 99).Return((*repo.Model)(nil), errors.New("not found"))

	err := service.Delete(99)
	assert.EqualError(t, err, "not found")

	mr.AssertExpectations(t)
}

func TestService_Delete_DeleteError(t *testing.T) {
	mr := new(MockRepo)
	service := svc.NewService(mr)

	now := time.Now().Truncate(time.Second)
	existing := &repo.Model{
		ID:          5,
		Title:       "X",
		Description: "Y",
		Price:       2.5,
		ImageURL:    "u",
		Quantity:    2,
		Category:    "beverages",
		CreatedAt:   now,
	}
	mr.On("GetByID", 5).Return(existing, nil)
	mr.On("Delete", 5).Return(errors.New("delete failed"))

	err := service.Delete(5)
	assert.EqualError(t, err, "delete failed")

	mr.AssertExpectations(t)
}
