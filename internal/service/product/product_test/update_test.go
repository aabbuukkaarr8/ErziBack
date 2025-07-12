package product

import (
	"errors"
	"testing"
	"time"

	repo "erzi_new/internal/repository/product"
	svc "erzi_new/internal/service/product"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_Update_Success(t *testing.T) {
	mr := new(MockRepo)
	service := svc.NewService(mr)
	now := time.Now()
	original := &repo.Product{
		ID:          7,
		Title:       "New Title",
		Description: "New Desc",
		Price:       4.56,
		ImageURL:    "http://new.img",
		Quantity:    20,
		Category:    "beverages",
		CreatedAt:   now,
	}

	newPrice := 2.2
	input := svc.UpdateProduct{
		ID:    original.ID,
		Price: &newPrice,
	}
	mr.On("GetByID", original.ID).Return(original, nil)
	mr.On("Update", mock.MatchedBy(func(p *repo.Product) bool {
		return p.ID == original.ID &&
			p.Price == newPrice &&
			p.Title == original.Title &&
			p.Description == original.Description &&
			p.ImageURL == original.ImageURL &&
			p.Quantity == original.Quantity &&
			p.Category == original.Category
	})).Return(original, nil)

	got, err := service.Update(input)
	assert.NoError(t, err)

	want := &svc.Product{}
	want.FillFromDB(original)
	assert.Equal(t, want, got)

	mr.AssertExpectations(t)
}

func TestService_Update_Fail(t *testing.T) {
	mr := new(MockRepo)
	service := svc.NewService(mr)
	now := time.Now()
	original := &repo.Product{
		ID:          2,
		Title:       "X",
		Description: "Y",
		Price:       4.56,
		ImageURL:    "http://new.img",
		Quantity:    20,
		Category:    "beverages",
		CreatedAt:   now,
	}
	newTitle := "New Title"
	input := svc.UpdateProduct{
		ID:    original.ID,
		Title: &newTitle,
	}
	mr.On("GetByID", original.ID).Return(original, nil)
	mr.On("Update", mock.MatchedBy(func(p *repo.Product) bool {
		return p.ID == original.ID &&
			p.Title == newTitle &&
			p.Description == original.Description &&
			p.Price == original.Price &&
			p.ImageURL == original.ImageURL &&
			p.Quantity == original.Quantity &&
			p.Category == original.Category
	})).Return((*repo.Product)(nil), errors.New("db fail"))

	got, err := service.Update(input)
	assert.Nil(t, got)
	assert.EqualError(t, err, "db fail")

	mr.AssertExpectations(t)
}
