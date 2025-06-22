package product

import "time"

type Product struct {
	ID          int
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
              RETURNING id, title, description, price, image_url, created_at, quantity`

	err := r.store.GetConn().QueryRow(
		query,
		p.Title,
		p.Description,
		p.Price,
		p.ImageURL,
		p.CreatedAt,
		p.Quantity,
	).Scan(&returnedP.ID, &returnedP.Title, &returnedP.Description, &returnedP.Price, &returnedP.ImageURL, &returnedP.CreatedAt, &returnedP.Quantity)
	if err != nil {
		return nil, err
	}
	return returnedP, nil
}
