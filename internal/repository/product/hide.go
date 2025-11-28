package product

import (
	"context"
)

func (r *Repository) Hide(ctx context.Context, id int, status bool) error {
	_, err := r.store.GetConn().QueryContext(ctx, `UPDATES products SET is_active = $1 WHERE id = $2`, status, id)
	if err != nil {
		return err
	}
	return nil
}
