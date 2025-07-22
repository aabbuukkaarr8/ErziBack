package product

import (
	"erzi_new/internal/repository/product"
	svc "erzi_new/internal/service/product"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestService_Create(t *testing.T) {
	mr := new(MockRepo)
	service := svc.NewService(mr)
	now := time.Now()
	input := svc.CreateProduct{
		Title:       "Water",
		Description: "Bottled spring water",
		Price:       1.23,
		Quantity:    42,
		Category:    "beverages",
	}

	repoProd := &product.Model{
		ID:          7,
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		ImageURL:    "",
		Quantity:    input.Quantity,
		Category:    input.Category,
		CreatedAt:   now,
	}
	mr.
		On("Create", mock.MatchedBy(func(p *product.Model) bool {
			return p.Title == input.Title &&
				p.Description == input.Description &&
				p.Price == input.Price &&
				p.Quantity == input.Quantity

		})).
		Return(repoProd, nil)

	out, err := service.Create(input)
	assert.NoError(t, err)
	assert.Equal(t, svc.Model{
		ID:          repoProd.ID,
		Title:       repoProd.Title,
		Description: repoProd.Description,
		Price:       repoProd.Price,
		ImageURL:    repoProd.ImageURL,
		Quantity:    repoProd.Quantity,
		Category:    repoProd.Category,
		CreatedAt:   repoProd.CreatedAt,
	}, *out)

	mr.AssertExpectations(t)
}
