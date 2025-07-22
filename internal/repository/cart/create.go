package cart

func (r *Repository) Create(userID int, status string) (*Model, error) {
	query := `INSERT INTO carts (user_id, status) VALUES ($1, $2) RETURNING id, user_id, status`
	row := r.store.GetConn().QueryRow(query, userID, status)

	var cart Model
	if err := row.Scan(&cart.ID, &cart.UserID, &cart.Status); err != nil {
		return nil, err
	}
	return &cart, nil
}
