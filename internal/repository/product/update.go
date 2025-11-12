package product

import "context"

func (r *Repository) Update(ctx context.Context, p *Model) (*Model, error) {
	updated := &Model{}
	query := `
    UPDATE products
    SET title = $1, description = $2, price = $3, image_url = $4, quantity = $5, category = $6
    WHERE id = $7
    RETURNING id, title, description, price, image_url, quantity, category, created_at
  `

	err := r.store.GetConn().QueryRowContext(ctx,
		query,
		p.Title,
		p.Description,
		p.Price,
		p.ImageURL,
		p.IsActive,
		p.Category,
		p.ID,
	).Scan(
		&updated.ID,
		&updated.Title,
		&updated.Description,
		&updated.Price,
		&updated.ImageURL,
		&updated.IsActive,
		&updated.Category,
		&updated.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return updated, nil
}
