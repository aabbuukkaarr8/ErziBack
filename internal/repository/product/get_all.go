package product

func (r *Repository) GetAllProducts() ([]Product, error) {
	var products []Product
	rows, err := r.store.GetConn().Query(`SELECT id, title, description, price, image_url, quantity, created_at FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.ImageURL, &p.Quantity, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
