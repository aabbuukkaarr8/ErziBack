package product

func (r *Repository) GetAllProducts() ([]Model, error) {
	var products []Model
	rows, err := r.store.GetConn().Query(`SELECT id, title, description, price, image_url, quantity, category, created_at FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Model{}
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.ImageURL, &p.Quantity, &p.Category, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
