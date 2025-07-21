package cart

//
//import (
//	"database/sql"
//	cartrepo "erzi_new/internal/repository/cart"
//	"erzi_new/internal/store"
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/stretchr/testify/require"
//	"testing"
//)
//
//func TestRepository_CreateCart(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	require.NoError(t, err)
//	defer db.Close()
//
//	s := store.New()
//	s.SetConn(db)
//	repo := cartrepo.NewRepository(s)
//
//	userID := 123
//	expectedID := 1
//
//	mock.ExpectQuery(`INSERT INTO carts \(user_id\) VALUES \(\$1\) RETURNING id, user_id`).
//		WithArgs(userID).
//		WillReturnRows(
//			sqlmock.NewRows([]string{"id", "user_id"}).AddRow(expectedID, userID),
//		)
//
//	cart, err := repo.Create(userID, status)
//	require.NoError(t, err)
//	require.NotNil(t, cart)
//	require.Equal(t, expectedID, cart.ID)
//	require.Equal(t, userID, cart.UserID)
//
//	require.NoError(t, mock.ExpectationsWereMet())
//}
//
//type mockStore struct {
//	db *sql.DB
//}
//
//func (m *mockStore) GetConn() *sql.DB {
//	return m.db
//}
