package cartItem

import "context"

func (r *Repository) GetByID(ctx context.Context, id int) (*Model, error) {
	const query = `
        SELECT id, cart_id, product_id, quantity, created_at
          FROM cart_items
         WHERE id = $1
    `
	row := r.store.GetConn().QueryRowContext(ctx, query, id)

	var itm Model
	if err := row.Scan(
		&itm.ID,
		&itm.CartID,
		&itm.ProductID,
		&itm.Quantity,
		&itm.CreatedAt,
	); err != nil {
		return nil, err
	}
	return &itm, nil
}
