// internal/repository/product/product_test/create_test.go
package product_test

import (
	"regexp"
	"testing"
	"time"

	productrepo "erzi_new/internal/repository/product"
	"erzi_new/internal/store"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := productrepo.NewRepository(s)

	now := time.Now().Truncate(time.Second)
	p := &productrepo.Model{
		Title:       "Product Title",
		Description: "Product Description",
		Price:       123,
		ImageURL:    "http://image.url",
		Category:    "beverages",
		CreatedAt:   now,
		Quantity:    1,
	}

	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO products (title, description, price, image_url, category, created_at, quantity)
         VALUES ($1,$2,$3,$4,$5,$6,$7)
         RETURNING id, title, description, price, image_url, category, created_at, quantity`,
	)).
		WithArgs(
			p.Title,
			p.Description,
			p.Price,
			p.ImageURL,
			p.Category,
			now,
			p.Quantity,
		).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "title", "description", "price", "image_url", "category", "created_at", "quantity",
		}).AddRow(
			42,
			p.Title,
			p.Description,
			p.Price,
			p.ImageURL,
			p.Category,
			now,
			p.Quantity,
		),
		)

	returned, err := repo.Create(p)
	assert.NoError(t, err)
	assert.Equal(t, 42, returned.ID)
	assert.Equal(t, p.Title, returned.Title)
	assert.Equal(t, p.Description, returned.Description)
	assert.Equal(t, p.Price, returned.Price)
	assert.Equal(t, p.ImageURL, returned.ImageURL)
	assert.Equal(t, p.Category, returned.Category)
	assert.Equal(t, now, returned.CreatedAt)
	assert.Equal(t, p.Quantity, returned.Quantity)

	assert.NoError(t, mock.ExpectationsWereMet())
}
