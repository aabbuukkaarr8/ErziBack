package user

import (
	repoUser "erzi_new/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (s *Service) Create(input CreateUser) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	toDB := repoUser.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Role:      input.Role,
		CreatedAt: time.Now(),
	}

	created, err := s.repo.Create(&toDB)
	if err != nil {
		return nil, err
	}

	fromDB := User{
		ID:        created.ID,
		Username:  created.Username,
		Email:     created.Email,
		Password:  created.Password,
		Role:      created.Role,
		CreatedAt: created.CreatedAt,
	}

	return &fromDB, nil
}
