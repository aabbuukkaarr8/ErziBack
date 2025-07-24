package user

import (
	"github.com/google/uuid"
	"time"
)

func (r *Repository) Create(u *User) (*User, error) {
	if u.Role == "" {
		u.Role = "user"
	}
	returnedU := &User{}
	u.CreatedAt = time.Now()
	query := `INSERT INTO users (id, username, email, password, role, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id, username, email, password, role, created_at`
	err := r.store.GetConn().QueryRow(
		query,
		uuid.New(),
		u.Username,
		u.Email,
		u.Password,
		u.Role,
		u.CreatedAt,
	).Scan(&returnedU.ID, &returnedU.Username, &returnedU.Email, &returnedU.Password, &returnedU.Role, &returnedU.CreatedAt)
	if err != nil {
		return nil, err
	}
	if returnedU.Role == "" {
		returnedU.Role = "user"
	}
	return returnedU, nil

}
