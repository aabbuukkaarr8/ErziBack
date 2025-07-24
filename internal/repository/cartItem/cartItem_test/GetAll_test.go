package cartitem

import (
	"errors"
	cartitemrepo "erzi_new/internal/repository/cartItem"
	"erzi_new/internal/store"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetAll_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	cartID := 7
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "cart_id", "product_id", "quantity", "created_at",
		"title", "price", "image_url",
	}).
		AddRow(1, cartID, 42, 2, now, "Water 5L", 120.50, "url1").
		AddRow(2, cartID, 43, 1, now, "Water 1.5L", 30.00, "url2")

	mock.ExpectQuery(`FROM\s+cart_items\s+JOIN\s+products`).
		WithArgs(cartID).
		WillReturnRows(rows)

	items, err := repo.GetAll(cartID)
	require.NoError(t, err)
	require.Len(t, items, 2)

	first := items[0]
	require.Equal(t, 1, first.ID)
	require.Equal(t, cartID, first.CartID)
	require.Equal(t, 42, first.ProductID)
	require.Equal(t, 2, first.Quantity)
	require.Equal(t, "Water 5L", first.Product.Title)
	require.Equal(t, 120.50, first.Product.Price)
	require.Equal(t, "url1", first.Product.ImageURL)

	second := items[1]
	require.Equal(t, 2, second.ID)
	require.Equal(t, cartID, second.CartID)
	require.Equal(t, 43, second.ProductID)
	require.Equal(t, 1, second.Quantity)
	require.Equal(t, "Water 1.5L", second.Product.Title)
	require.Equal(t, 30.00, second.Product.Price)
	require.Equal(t, "url2", second.Product.ImageURL)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetAll_Empty(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	cartID := 99

	mock.ExpectQuery(`FROM\s+cart_items\s+JOIN\s+products`).
		WithArgs(cartID).
		WillReturnRows(sqlmock.NewRows(
			[]string{"id", "cart_id", "product_id", "quantity", "created_at", "title", "price", "image_url"},
		))

	items, err := repo.GetAll(cartID)
	require.NoError(t, err)
	require.Empty(t, items)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetAll_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	cartID := 5

	mock.ExpectQuery(`FROM\s+cart_items\s+JOIN\s+products`).
		WithArgs(cartID).
		WillReturnError(errors.New("db error"))

	items, err := repo.GetAll(cartID)
	require.Error(t, err)
	require.Contains(t, err.Error(), "db error")
	require.Nil(t, items)

	require.NoError(t, mock.ExpectationsWereMet())
}
