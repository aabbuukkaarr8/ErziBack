package product

import (
	repoUser "erzi_new/internal/repository/user"
	"time"
)

type CreateUser struct {
	Username string
	Email    string
	Password string
	Role     string
}

func (s *Service) Create(input CreateUser) (*repoUser.User, error) {
	toDB := repoUser.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		Role:      input.Role,
		CreatedAt: time.Now(),
	}
	created, err := s.repo.Create(&toDB)
	if err != nil {
		return nil, err
	}
	fromDB := repoUser.User{
		Username:  created.Username,
		Email:     created.Email,
		Password:  created.Password,
		Role:      created.Role,
		CreatedAt: time.Now(),
	}
	return &fromDB, nil
}
