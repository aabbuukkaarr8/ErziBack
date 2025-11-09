package cartItem

import "context"

func (r *Repository) GetAll(ctx context.Context, cartID int) ([]ModelWithProduct, error) {
	var out []ModelWithProduct
	query := `SELECT cart_items.id,
	cart_items.cart_id,
	cart_items.product_id,
	cart_items.quantity,
	cart_items.created_at,
	products.title,
	products.price,
	products.image_url,
	products.bulk_discount_quantity,
	products.bulk_discount_price
	FROM cart_items
	JOIN products
         	 ON cart_items.product_id = products.id
        	WHERE cart_items.cart_id = $1 
        	`
	rows, err := r.store.GetConn().QueryContext(ctx, query, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var itm ModelWithProduct
		if err := rows.Scan(
			&itm.ID,
			&itm.CartID,
			&itm.ProductID,
			&itm.Quantity,
			&itm.CreatedAt,
			&itm.Product.Title,
			&itm.Product.Price,
			&itm.Product.ImageURL,
			&itm.Product.BulkDiscountQuantity,
			&itm.Product.BulkDiscountPrice,
		); err != nil {
			return nil, err
		}
		out = append(out, itm)
	}
	return out, nil
}
