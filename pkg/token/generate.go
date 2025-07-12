package token

import (
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var jwtSecret = []byte("your_secret_key")

func GenerateJWT(userID int, role string, email string) (string, error) {
	tok := jwt.New()
	_ = tok.Set("user_id", userID)
	_ = tok.Set("role", role)
	_ = tok.Set("email", email)
	_ = tok.Set("exp", time.Now().Add(24*time.Hour).Unix())
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.HS256, jwtSecret))
	if err != nil {
		return "", err

	}
	return string(signed), nil
}
