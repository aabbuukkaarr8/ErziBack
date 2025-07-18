package cart

func (r *Repository) CreateCart(userID int) (*Cart, error) {
	query := `INSERT INTO carts (user_id) VALUES ($1) RETURNING id, user_id`
	row := r.store.GetConn().QueryRow(query, userID)

	var cart Cart
	if err := row.Scan(&cart.ID, &cart.UserID); err != nil {
		return nil, err
	}
	return &cart, nil
}
