package product

func (r *Repository) Update(p *Product) (*Product, error) {
	updated := &Product{}
	query := `
    UPDATE products
    SET title = $1, description = $2, price = $3, image_url = $4, quantity = $5
    WHERE id = $6
    RETURNING id, title, description, price, image_url, quantity, created_at
  `

	err := r.store.GetConn().QueryRow(
		query,
		p.Title,
		p.Description,
		p.Price,
		p.ImageURL,
		p.Quantity,
		p.ID,
	).Scan(
		&updated.ID,
		&updated.Title,
		&updated.Description,
		&updated.Price,
		&updated.ImageURL,
		&updated.Quantity,
		&updated.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return updated, nil
}
