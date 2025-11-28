package product

import (
	productrepo "erzi_new/internal/repository/product"
	"erzi_new/internal/store"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestRepository_Delete_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := productrepo.NewRepository(s)

	mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM products WHERE id = $1`,
	)).
		WithArgs(7).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.Delete(7)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepository_Delete_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := productrepo.NewRepository(s)

	mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM products WHERE id = $1`,
	)).
		WithArgs(999).
		WillReturnResult(sqlmock.NewResult(0, 0)) // no rows deleted

	err = repo.Delete(999)
	assert.EqualError(t, err, "product not found")
	assert.NoError(t, mock.ExpectationsWereMet())
}
