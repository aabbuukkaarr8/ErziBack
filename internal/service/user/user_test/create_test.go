package user

import (
	repo "erzi_new/internal/repository/user"
	"erzi_new/internal/service/user"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestService_Create(t *testing.T) {
	mockRepo := new(MockUserRepo)
	service := user.NewService(mockRepo)

	input := user.CreateUser{
		Username: "kuro",
		Email:    "kuro@example.com",
		Password: "hashedpass",
		Role:     "user",
	}

	expected := &repo.User{
		ID:        1,
		Username:  "kuro",
		Email:     "kuro@example.com",
		Password:  "hashedpass",
		Role:      "user",
		CreatedAt: time.Now(),
	}

	mockRepo.On("Create", mock.AnythingOfType("*user.User")).Return(expected, nil)

	actual, err := service.Create(input)
	require.NoError(t, err)
	require.Equal(t, expected.ID, actual.ID)
	require.Equal(t, expected.Username, actual.Username)
}
