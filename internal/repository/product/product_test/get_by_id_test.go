package product

import (
	"database/sql"
	product2 "erzi_new/internal/repository/product"
	"erzi_new/internal/store"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
	"time"
)

func TestRepository_GetByID_Succes(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()
	s := store.New()
	s.SetConn(db)
	repo := product2.NewRepository(s)
	cols := []string{"id", "title", "description", "price", "image_url", "quantity", "category", "created_at"}
	createdAt := time.Now()
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT id, title, description, price, image_url, quantity, category, created_at FROM products WHERE id = $1`,
	)).
		WithArgs(7).
		WillReturnRows(sqlmock.NewRows(cols).
			AddRow(7, "Tasty Water", "Pure spring", 2.50, "", 100, "waterBall", createdAt),
		)
	p, err := repo.GetByID(7)
	assert.NoError(t, err)
	assert.Equal(t, 7, p.ID)
	assert.Equal(t, "Tasty Water", p.Title)
	assert.Equal(t, "Pure spring", p.Description)
	assert.Equal(t, 2.50, p.Price)
	assert.Equal(t, "", p.ImageURL)
	assert.Equal(t, 100, p.Quantity)
	assert.Equal(t, "waterBall", p.Category)
	assert.Equal(t, createdAt, p.CreatedAt)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()
	s := store.New()
	s.SetConn(db)
	repo := product2.NewRepository(s)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT id, title, description, price, image_url, quantity, category, created_at FROM products WHERE id = $1`,
	)).
		WithArgs(99).
		WillReturnError(sql.ErrNoRows)
	p, err := repo.GetByID(99)
	assert.Nil(t, p)
	assert.EqualError(t, err, "sql: no rows in result set")

	assert.NoError(t, mock.ExpectationsWereMet())

}
