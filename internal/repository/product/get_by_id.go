package product

import "context"

func (r *Repository) GetByID(ctx context.Context, id int) (*Model, error) {
	p := Model{}
	err := r.store.GetConn().
		QueryRowContext(ctx, `SELECT id, title, description, price, image_url, quantity, category, created_at FROM products WHERE id = $1`, id).
		Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.ImageURL, &p.Quantity, &p.Category, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
