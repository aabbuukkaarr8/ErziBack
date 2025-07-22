package product

import (
	"database/sql"
	"erzi_new/internal/repository/product"
	"erzi_new/internal/store"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestRepository_Update_Succes(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()
	s := store.New()
	s.SetConn(db)
	repo := product.NewRepository(s)
	now := time.Now()
	input := &product.Model{
		ID:          7,
		Title:       "New Title",
		Description: "New Desc",
		Price:       9.99,
		ImageURL:    "http://new.img",
		Quantity:    42,
		Category:    "WaterBall",
	}
	query := `UPDATE products SET title = $1, description = $2, price = $3, image_url = $4, quantity = $5, category = $6 WHERE id = $7 RETURNING id, title, description, price, image_url, quantity, category, created_at`
	cols := []string{"id", "title", "description", "price", "image_url", "quantity", "category", "created_at"}
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(input.Title, input.Description, input.Price, input.ImageURL, input.Quantity, input.Category, input.ID).
		WillReturnRows(sqlmock.NewRows(cols).
			AddRow(input.ID, input.Title, input.Description, input.Price, input.ImageURL, input.Quantity, input.Category, now),
		)
	updated, err := repo.Update(input)
	assert.NoError(t, err)
	assert.Equal(t, input.ID, updated.ID)
	assert.Equal(t, input.Title, updated.Title)
	assert.Equal(t, input.Description, updated.Description)
	assert.Equal(t, input.Price, updated.Price)
	assert.Equal(t, input.ImageURL, updated.ImageURL)
	assert.Equal(t, input.Quantity, updated.Quantity)
	assert.Equal(t, input.Category, updated.Category)
	assert.Equal(t, now, updated.CreatedAt)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestRepository_Update_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := product.NewRepository(s)

	input := &product.Model{
		ID:          999,
		Title:       "X",
		Description: "Y",
		Price:       0.1,
		ImageURL:    "http://img",
		Category:    "WaterBall",
		Quantity:    1,
	}

	query := `UPDATE products SET title = $1, description = $2, price = $3, image_url = $4, quantity = $5, category = $6 WHERE id = $7 RETURNING id, title, description, price, image_url, quantity, category, created_at`
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(input.Title, input.Description, input.Price, input.ImageURL, input.Quantity, input.Category, input.ID).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.Update(input)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Update_OnlyPrice(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := product.NewRepository(s)

	now := time.Now().Truncate(time.Second)
	input := &product.Model{
		ID:    7,
		Price: 9.99,
	}

	query := `UPDATE products SET title = $1, description = $2, price = $3, image_url = $4, quantity = $5, category = $6 WHERE id = $7 RETURNING id, title, description, price, image_url, quantity, category, created_at`
	cols := []string{"id", "title", "description", "price", "image_url", "quantity", "category", "created_at"}
	mock.ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs("", "", input.Price, "", 0, "", input.ID).
		WillReturnRows(sqlmock.NewRows(cols).
			AddRow(input.ID, "", "", input.Price, "", 0, "", now),
		)

	updated, err := repo.Update(input)
	assert.NoError(t, err)
	assert.Equal(t, input.ID, updated.ID)
	assert.Equal(t, "", updated.Title)
	assert.Equal(t, "", updated.Description)
	assert.Equal(t, input.Price, updated.Price)
	assert.Equal(t, "", updated.ImageURL)
	assert.Equal(t, 0, updated.Quantity)
	assert.Equal(t, "", updated.Category)
	assert.Equal(t, now, updated.CreatedAt)

	assert.NoError(t, mock.ExpectationsWereMet())
}
