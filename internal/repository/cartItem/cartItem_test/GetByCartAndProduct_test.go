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

func TestRepository_GetByCartAndProduct_Found(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	cartID := 10
	productID := 42
	now := time.Now()

	rows := sqlmock.NewRows([]string{"id", "cart_id", "product_id", "quantity", "created_at"}).
		AddRow(1, cartID, productID, 3, now)

	mock.ExpectQuery(`SELECT\s+id,\s*cart_id,\s*product_id,\s*quantity,\s*created_at`).
		WithArgs(cartID, productID).
		WillReturnRows(rows)

	itm, err := repo.GetByCartAndProduct(cartID, productID)
	require.NoError(t, err)
	require.NotNil(t, itm)
	require.Equal(t, 1, itm.ID)
	require.Equal(t, cartID, itm.CartID)
	require.Equal(t, productID, itm.ProductID)
	require.Equal(t, 3, itm.Quantity)
	require.WithinDuration(t, now, itm.CreatedAt, time.Second)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetByCartAndProduct_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	cartID := 99
	productID := 100

	mock.ExpectQuery(`SELECT\s+id,\s*cart_id,\s*product_id,\s*quantity,\s*created_at`).
		WithArgs(cartID, productID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "cart_id", "product_id", "quantity", "created_at"}))

	itm, err := repo.GetByCartAndProduct(cartID, productID)
	require.NoError(t, err)
	require.Nil(t, itm)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetByCartAndProduct_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	cartID := 5
	productID := 6

	mock.ExpectQuery(`SELECT\s+id,\s*cart_id,\s*product_id,\s*quantity,\s*created_at`).
		WithArgs(cartID, productID).
		WillReturnError(errors.New("db error"))

	itm, err := repo.GetByCartAndProduct(cartID, productID)
	require.Error(t, err)
	require.Contains(t, err.Error(), "db error")
	require.Nil(t, itm)

	require.NoError(t, mock.ExpectationsWereMet())
}
