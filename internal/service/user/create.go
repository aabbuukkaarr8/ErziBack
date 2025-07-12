package user

import (
	repoUser "erzi_new/internal/repository/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CreateUser struct {
	Username string
	Email    string
	Password string
	Role     string
}

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}

func (s *Service) Create(input CreateUser) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	toDB := repoUser.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  string(hashedPassword), // Хеш сохраняется в БД
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
