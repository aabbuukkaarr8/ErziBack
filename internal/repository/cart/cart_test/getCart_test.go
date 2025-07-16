package cart

import (
	cartrepo "erzi_new/internal/repository/cart"
	"erzi_new/internal/store"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRepository_GetCart(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartrepo.NewRepository(s)

	userID := 42
	expectedCartID := 10

	mock.ExpectQuery(`SELECT id FROM carts WHERE user_id = \$1`).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedCartID))

	id, err := repo.GetCart(userID)
	require.NoError(t, err)
	require.Equal(t, expectedCartID, id)

	require.NoError(t, mock.ExpectationsWereMet())
}
