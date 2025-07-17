package cartItem

func (r *Repository) GetByID(id int) (*CartItem, error) {
	const query = `
        SELECT id, cart_id, product_id, quantity, created_at
          FROM cart_items
         WHERE id = $1
    `
	row := r.store.GetConn().QueryRow(query, id)

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
