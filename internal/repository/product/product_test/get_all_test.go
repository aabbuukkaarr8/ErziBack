package product

import (
	product2 "erzi_new/internal/repository/product"
	"erzi_new/internal/store"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestRepository_GetAllProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()
	s := store.New()
	s.SetConn(db)
	repo := product2.NewRepository(s)
	now := time.Now()
	cols := []string{"id", "title", "description", "price", "image_url", "quantity", "category", "created_at"}
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT id, title, description, price, image_url, quantity, category, created_at FROM products`,
	)).
		WillReturnRows(sqlmock.NewRows(cols).
			AddRow(1, "Water", "Spring water", 1.23, "http://img1", 10, "waterBall", now).
			AddRow(2, "Juice", "Orange juice", 2.34, "http://img2", 5, "waterBall", now),
		)

	list, err := repo.GetAllProducts()
	assert.NoError(t, err)
	assert.Len(t, list, 2)

	assert.Equal(t, 1, list[0].ID)
	assert.Equal(t, "Water", list[0].Title)
	assert.Equal(t, "Spring water", list[0].Description)
	assert.Equal(t, 1.23, list[0].Price)
	assert.Equal(t, "http://img1", list[0].ImageURL)
	assert.Equal(t, 10, list[0].Quantity)
	assert.Equal(t, "waterBall", list[0].Category)
	assert.Equal(t, now, list[0].CreatedAt)

	assert.Equal(t, 2, list[1].ID)
	assert.Equal(t, "Juice", list[1].Title)
	assert.Equal(t, "Orange juice", list[1].Description)
	assert.Equal(t, 2.34, list[1].Price)
	assert.Equal(t, "http://img2", list[1].ImageURL)
	assert.Equal(t, 5, list[1].Quantity)
	assert.Equal(t, "waterBall", list[1].Category)
	assert.Equal(t, now, list[1].CreatedAt)
	assert.NoError(t, mock.ExpectationsWereMet())

}
