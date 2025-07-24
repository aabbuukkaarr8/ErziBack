package cartitem

import (
	"database/sql"
	cartitemrepo "erzi_new/internal/repository/cartItem"
	"erzi_new/internal/store"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRepository_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	cartID := 1
	productID := 2

	mock.ExpectQuery(`INSERT INTO cart_items \(cart_id, product_id, quantity\) VALUES \(\$1, \$2, 1\) RETURNING id, cart_id, product_id, quantity, created_at`).
		WithArgs(cartID, productID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "cart_id", "product_id", "quantity", "created_at"}).
			AddRow(10, cartID, productID, 1, time.Now()))

	item, err := repo.Create(cartID, productID)
	require.NoError(t, err)
	require.NotNil(t, item)
	require.Equal(t, cartID, item.CartID)
	require.Equal(t, productID, item.ProductID)
	require.Equal(t, 1, item.Quantity)

	require.NoError(t, mock.ExpectationsWereMet())
}

type mockStore struct {
	db *sql.DB
}

func (m *mockStore) GetConn() *sql.DB {
	return m.db
}
