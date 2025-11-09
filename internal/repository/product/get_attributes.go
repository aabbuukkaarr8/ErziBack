package product

import "context"

func (r *Repository) GetAttributes(ctx context.Context, id int) ([]Attribute, error) {
	var a []Attribute
	rows, err := r.store.GetConn().QueryContext(ctx, `SELECT * FROM product_attributes WHERE product_id = $1`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		m := Attribute{}
		err := rows.Scan(&m.ID, &m.ProductID, &m.Key, &m.Value)
		if err != nil {
			return nil, err
		}
		a = append(a, m)

	}
	return a, nil
}
