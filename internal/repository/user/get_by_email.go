package user

import "context"

func (r *Repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	row := r.store.GetConn().QueryRowContext(ctx, `SELECT * FROM users WHERE email = $1`, email)
	u := &User{}
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.Password, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}
