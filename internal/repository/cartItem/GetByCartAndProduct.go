package cartItem

import "database/sql"

func (r *Repository) GetByCartAndProduct(cartID, productID int) (*CartItem, error) {
	const query = `
        SELECT id, cart_id, product_id, quantity, created_at
          FROM cart_items
         WHERE cart_id = $1 AND product_id = $2`
	row := r.store.GetConn().QueryRow(query, cartID, productID)

	var itm CartItem
	if err := row.Scan(
		&itm.ID, &itm.CartID, &itm.ProductID, &itm.Quantity, &itm.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &itm, nil
}
