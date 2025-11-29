package product

func (r *Repository) GetAllProducts() ([]Model, error) {
	var products []Model
	rows, err := r.store.GetConn().Query(`SELECT id, title, description, image_url, is_active, category, created_at, prices FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := Model{}
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.ImageURL, &p.IsActive, &p.Category, &p.CreatedAt, &p.Prices)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}
