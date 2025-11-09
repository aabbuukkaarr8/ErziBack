package cartItem

import "context"

func (r *Repository) DeleteAll(ctx context.Context, cartID int) error {
	query := "UPDATE carts SET status = 'deleted' WHERE id = $1"
	_, err := r.store.GetConn().ExecContext(ctx, query, cartID)
	return err

}
