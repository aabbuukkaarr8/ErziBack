package product

import "context"

func (r *Repository) GetByID(ctx context.Context, id int) (*Model, error) {
	p := Model{}
	err := r.store.GetConn().
		QueryRowContext(ctx, `SELECT id, title, description, image_url, is_active, category, created_at, prices FROM products WHERE id = $1`, id).
		Scan(&p.ID, &p.Title, &p.Description, &p.ImageURL, &p.IsActive, &p.Category, &p.CreatedAt, &p.Prices)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
