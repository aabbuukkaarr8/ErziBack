package cartItem

func (r *Repository) Create(cartID, productID int) (*Model, error) {
	query := `INSERT INTO cart_items (cart_id, product_id, quantity) VALUES ($1, $2, 1) RETURNING id, cart_id, product_id, quantity, created_at`
	var cartItem Model

	err := r.store.GetConn().QueryRow(query, cartID, productID).Scan(
		&cartItem.ID,
		&cartItem.CartID,
		&cartItem.ProductID,
		&cartItem.Quantity,
		&cartItem.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}
