package product

import "context"

func (r *Repository) Create(ctx context.Context, p *Model) (*Model, error) {
	returnedP := &Model{}
	query := `INSERT INTO products (title, description, image_url, category, created_at, is_active, prices)
              VALUES ($1,$2,$3,$4,$5,$6,$7)
              RETURNING id, title, description, image_url, category, created_at, is_active, prices`

	err := r.store.GetConn().QueryRowContext(ctx,
		query,
		p.Title,
		p.Description,
		p.ImageURL,
		p.Category,
		p.CreatedAt,
		p.IsActive,
		p.Prices,
	).Scan(&returnedP.ID, &returnedP.Title, &returnedP.Description, &returnedP.ImageURL, &returnedP.Category, &returnedP.CreatedAt, &returnedP.IsActive, &returnedP.Prices)
	if err != nil {
		return nil, err
	}
	return returnedP, nil
}
