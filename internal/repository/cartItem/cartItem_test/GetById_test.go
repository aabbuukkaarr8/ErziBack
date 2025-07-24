package cartitem

import (
	"database/sql"
	"errors"
	cartitemrepo "erzi_new/internal/repository/cartItem"
	"erzi_new/internal/store"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestRepository_GetByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	itemID := 5
	cartID := 10
	productID := 42
	qty := 3
	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"id", "cart_id", "product_id", "quantity", "created_at",
	}).AddRow(itemID, cartID, productID, qty, now)

	mock.ExpectQuery(`SELECT\s+id,\s*cart_id,\s*product_id,\s*quantity,\s*created_at`).
		WithArgs(itemID).
		WillReturnRows(rows)

	itm, err := repo.GetByID(itemID)
	require.NoError(t, err)
	require.NotNil(t, itm)
	require.Equal(t, itemID, itm.ID)
	require.Equal(t, cartID, itm.CartID)
	require.Equal(t, productID, itm.ProductID)
	require.Equal(t, qty, itm.Quantity)
	require.WithinDuration(t, now, itm.CreatedAt, time.Second)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	itemID := 99

	mock.ExpectQuery(`SELECT\s+id,\s*cart_id,\s*product_id,\s*quantity,\s*created_at`).
		WithArgs(itemID).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "cart_id", "product_id", "quantity", "created_at",
		}))

	itm, err := repo.GetByID(itemID)
	require.Error(t, err)
	require.Nil(t, itm)
	require.Equal(t, sql.ErrNoRows, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_GetByID_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	itemID := 7

	mock.ExpectQuery(`SELECT\s+id,\s*cart_id,\s*product_id,\s*quantity,\s*created_at`).
		WithArgs(itemID).
		WillReturnError(errors.New("db error"))

	itm, err := repo.GetByID(itemID)
	require.Error(t, err)
	require.Nil(t, itm)
	require.Contains(t, err.Error(), "db error")

	require.NoError(t, mock.ExpectationsWereMet())
}
