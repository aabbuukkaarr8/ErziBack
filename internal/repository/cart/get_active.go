package cart

func (r *Repository) GetActive(userID int) (int, error) {
	var id int
	query := `SELECT id FROM carts WHERE user_id = $1 AND status = 'active'`
	err := r.store.GetConn().QueryRow(query, userID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
