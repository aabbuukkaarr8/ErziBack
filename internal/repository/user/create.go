package cart

import (
	"time"
)

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}

func (r *Repository) Create(u *User) (*User, error) {
	returnedU := &User{}
	u.CreatedAt = time.Now()
	query := `INSERT INTO users (username, email, password, role, created_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, username, email, password, role, created_at`
	err := r.store.GetConn().QueryRow(query,
		u.Username,
		u.Email,
		u.Password,
		u.Role,
		u.CreatedAt,
	).Scan(&returnedU.ID, &returnedU.Username, &returnedU.Email, &returnedU.Password, &returnedU.Role, &returnedU.CreatedAt)
	if err != nil {
		return nil, err
	}
	return returnedU, nil

}
