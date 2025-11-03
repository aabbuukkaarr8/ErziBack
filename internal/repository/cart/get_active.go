package cart

import (
	"context"
	"github.com/google/uuid"
)

func (r *Repository) GetActive(ctx context.Context, userID uuid.UUID) (int, error) {
	var id int
	query := `SELECT id FROM carts WHERE user_id = $1 AND status = 'active'`
	err := r.store.GetConn().QueryRowContext(ctx, query, userID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
