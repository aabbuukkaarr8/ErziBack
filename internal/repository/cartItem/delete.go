package cartItem

import "context"

func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.store.GetConn().ExecContext(ctx,
		`DELETE FROM cart_items WHERE id = $1`, id,
	)
	return err
}
