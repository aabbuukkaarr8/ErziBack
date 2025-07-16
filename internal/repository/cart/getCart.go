package cart

func (r *Repository) GetCart(userID int) (int, error) {
	var id int
	query := `SELECT id FROM carts WHERE user_id = $1`
	err := r.store.GetConn().QueryRow(query, userID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
