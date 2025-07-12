package user_test

import (
	"erzi_new/internal/repository/user"
	"erzi_new/internal/store"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRepository_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	s := store.New()
	s.SetConn(db)
	repo := user.NewRepository(s)

	mock.ExpectQuery(`INSERT INTO users`).
		WithArgs("kuro", "kuro@example.com", "hashedpass", "user", sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{
			"id", "username", "email", "password", "role", "created_at",
		}).AddRow(1, "kuro", "kuro@example.com", "hashedpass", "user", time.Now()))

	u := &user.User{
		Username:  "kuro",
		Email:     "kuro@example.com",
		Password:  "hashedpass",
		Role:      "user",
		CreatedAt: time.Now(),
	}

	created, err := repo.Create(u)
	require.NoError(t, err)
	require.Equal(t, 1, created.ID)
	require.Equal(t, "kuro", created.Username)
}
