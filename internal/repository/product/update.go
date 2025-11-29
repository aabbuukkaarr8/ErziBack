package product

import "context"

func (r *Repository) Update(ctx context.Context, p *Model) (*Model, error) {
	updated := &Model{}
	query := `
    UPDATE products
    SET title = $1, description = $2, image_url = $3, is_active = $4, category = $5, prices = $6
    WHERE id = $7
    RETURNING id, title, description, image_url, is_active, category, created_at, prices
  `

	err := r.store.GetConn().QueryRowContext(ctx,
		query,
		p.Title,
		p.Description,
		p.ImageURL,
		p.IsActive,
		p.Category,
		p.Prices,
		p.ID,
	).Scan(
		&updated.ID,
		&updated.Title,
		&updated.Description,
		&updated.ImageURL,
		&updated.IsActive,
		&updated.Category,
		&updated.CreatedAt,
		&updated.Prices,
	)

	if err != nil {
		return nil, err
	}
	return updated, nil
}
