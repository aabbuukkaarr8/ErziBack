package cart

import "github.com/google/uuid"

func (r *Repository) GetDelete(userID uuid.UUID) (int, error) {
	var id int
	query := `SELECT id FROM carts WHERE user_id = $1 AND status = 'deleted'`
	err := r.store.GetConn().QueryRow(query, userID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
