package product

func (r *Repository) Create(p *Model) (*Model, error) {
	returnedP := &Model{}
	query := `INSERT INTO products (title, description, price, image_url, category, created_at, quantity, bulk_discount_quantity, bulk_discount_price)
              VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
              RETURNING id, title, description, price, image_url, category, created_at, quantity, bulk_discount_quantity, bulk_discount_price`

	err := r.store.GetConn().QueryRow(
		query,
		p.Title,
		p.Description,
		p.Price,
		p.ImageURL,
		p.Category,
		p.CreatedAt,
		p.Quantity,
		p.BulkDiscountQuantity,
		p.BulkDiscountPrice,
	).Scan(&returnedP.ID, &returnedP.Title, &returnedP.Description, &returnedP.Price, &returnedP.ImageURL, &returnedP.Category, &returnedP.CreatedAt, &returnedP.Quantity, &returnedP.BulkDiscountQuantity, &returnedP.BulkDiscountPrice)
	if err != nil {
		return nil, err
	}
	return returnedP, nil
}
