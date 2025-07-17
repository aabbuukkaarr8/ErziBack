package user

import (
	"errors"
	"erzi_new/pkg/token"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(email, password string) (string, error) {
	u, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	tok, err := token.GenerateJWT(u.ID, u.Role, u.Email)
	if err != nil {
		return "", err
	}

	return tok, nil
}
