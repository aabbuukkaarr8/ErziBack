package cartitem

import (
	"errors"
	cartitemrepo "erzi_new/internal/repository/cartItem"
	"erzi_new/internal/store"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestRepository_Delete_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	mock.ExpectExec(`DELETE FROM cart_items WHERE id = \$1`).
		WithArgs(5).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Delete(5)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Delete_ExecError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := cartitemrepo.NewRepository(s)

	mock.ExpectExec(`DELETE FROM cart_items WHERE id = \$1`).
		WithArgs(5).
		WillReturnError(errors.New("boom"))

	err = repo.Delete(5)
	require.Error(t, err)
	require.Contains(t, err.Error(), "boom")
	require.NoError(t, mock.ExpectationsWereMet())
}
