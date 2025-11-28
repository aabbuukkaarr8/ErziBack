package product

import (
	"context"
	"fmt"
)

func (r *Repository) GetStatus(ctx context.Context, id int) (bool, error) {
	var status bool
	err := r.store.GetConn().QueryRowContext(ctx, `SELECT is_active FROM products WHERE id = $1`, id).Scan(status)
	if err != nil {
		return false, fmt.Errorf("product.GetStatus: %w", err)
	}
	return status, nil
}
