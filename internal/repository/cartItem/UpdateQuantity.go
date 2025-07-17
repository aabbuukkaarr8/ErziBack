package cartItem

func (r *Repository) UpdateQuantity(id, quantity int) (*CartItem, error) {
	const query = `
        UPDATE cart_items
           SET quantity = $2
         WHERE id = $1
     RETURNING id, cart_id, product_id, quantity, created_at
    `
	row := r.store.GetConn().QueryRow(query, id, quantity)
	var itm CartItem
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
