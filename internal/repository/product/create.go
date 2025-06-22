package product

import "time"

type Product struct {
	ProductID   int
	Title       string
	Description string
	Price       float64
	ImageURL    string
	Quantity    int
	CreatedAt   time.Time
}

func (r *Repository) Create(p *Product) (*Product, error) {
	returnedP := &Product{}
	query := `INSERT INTO products (title, description, price, image_url, created_at, quantity)
              VALUES ($1, $2, $3, $4, $5, $6)
              RETURNING product_id, title, description, price, image_url, created_at, quantity`

	err := r.store.GetConn().QueryRow(
		query,
		p.Title,
		p.Description,
		p.Price,
		p.ImageURL,
		p.CreatedAt,
		p.Quantity,
	).Scan(&returnedP)
	if err != nil {
		return nil, err
	}
	return returnedP, nil
}
