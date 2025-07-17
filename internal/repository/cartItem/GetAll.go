package cartItem

import "time"

type CartItemWithProduct struct {
	ID        int
	CartID    int
	ProductID int
	Quantity  int
	CreatedAt time.Time
	Product   struct {
		Title    string
		Price    float64
		ImageURL string
	}
}

func (r *Repository) GetAll(cartID int) ([]*CartItemWithProduct, error) {
	query := `SELECT cart_items.id,
	cart_items.cart_id,
	cart_items.product_id,
	cart_items.quantity,
	cart_items.created_at,
	products.title,
	products.price,
	products.image_url
	FROM cart_items
	JOIN products
         	 ON cart_items.product_id = products.id
        	WHERE cart_items.cart_id = $1 
        	`
	rows, err := r.store.GetConn().Query(query, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []*CartItemWithProduct
	for rows.Next() {
		var itm CartItemWithProduct
		if err := rows.Scan(
			&itm.ID,
			&itm.CartID,
			&itm.ProductID,
			&itm.Quantity,
			&itm.CreatedAt,
			&itm.Product.Title,
			&itm.Product.Price,
			&itm.Product.ImageURL,
		); err != nil {
			return nil, err
		}
		out = append(out, &itm)
	}
	return out, nil
}
