package product

func (r *Repository) GetActiveProducts() ([]Model, error) {
	var products []Model
	rows, err := r.store.GetConn().Query(`SELECT id, title, description, price, image_url, is_active, category, created_at, bulk_discount_quantity, bulk_discount_price FROM products WHERE is_active = TRUE`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Model{}
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.ImageURL, &p.IsActive, &p.Category, &p.CreatedAt, &p.BulkDiscountQuantity, &p.BulkDiscountPrice)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
