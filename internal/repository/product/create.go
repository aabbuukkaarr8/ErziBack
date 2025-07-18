package product

func (r *Repository) Create(p *Product) (*Product, error) {
	returnedP := &Product{}
	query := `INSERT INTO products (title, description, price, image_url, category, created_at, quantity)
              VALUES ($1,$2,$3,$4,$5,$6,$7)
              RETURNING id, title, description, price, image_url, category, created_at, quantity`

	err := r.store.GetConn().QueryRow(
		query,
		p.Title,
		p.Description,
		p.Price,
		p.ImageURL,
		p.Category,
		p.CreatedAt,
		p.Quantity,
	).Scan(&returnedP.ID, &returnedP.Title, &returnedP.Description, &returnedP.Price, &returnedP.ImageURL, &returnedP.Category, &returnedP.CreatedAt, &returnedP.Quantity)
	if err != nil {
		return nil, err
	}
	return returnedP, nil
}
