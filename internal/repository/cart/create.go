package cart

type Cart struct {
	ID     int
	UserID int
}

func (r *Repository) Create(userID int) (*Cart, error) {
	var c Cart
	err := r.store.GetConn().QueryRow(`INSERT INTO carts(user_id) VALUES($1) RETURNING id, user_id`,
		userID,
	).Scan(&c.ID, &c.UserID)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
