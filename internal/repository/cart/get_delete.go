package cart

import (
	"context"
	"github.com/google/uuid"
)

func (r *Repository) RestoreCart(ctx context.Context, userID uuid.UUID) error {
	query := `UPDATE carts SET status = 'active' WHERE user_id = $1 AND status = 'deleted'`
	_, err := r.store.GetConn().ExecContext(ctx, query, userID)
	if err != nil {
		return err
	}
	return nil
}
