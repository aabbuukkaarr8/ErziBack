package product

func (r *Repository) GetByID(id int) (*Product, error) {
	p := Product{}
	err := r.store.GetConn().
		QueryRow(`SELECT id, title, description, price, image_url, quantity, created_at FROM products WHERE id = $1`, id).
		Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.ImageURL, &p.Quantity, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
