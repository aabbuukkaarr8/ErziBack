package product

func (r *Repository) CreateAttributes(a *Attribute) (*Attribute, error) {
	const query = `
		INSERT INTO product_attributes (product_id, key, value)
		VALUES ($1, $2, $3)
		RETURNING id, product_id, key, value
		`
	result := &Attribute{}
	err := r.store.GetConn().QueryRow(
		query,
		a.ProductID,
		a.Key,
		a.Value,
	).Scan(
		&result.ID,
		&result.ProductID,
		&result.Key,
		&result.Value,
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}
