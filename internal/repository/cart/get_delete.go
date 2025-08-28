package cart

import "github.com/google/uuid"

func (r *Repository) RestoreCart(userID uuid.UUID) error {
	query := `UPDATE carts SET status = 'active' WHERE user_id = $1 AND status = 'deleted'`
	_, err := r.store.GetConn().Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}
