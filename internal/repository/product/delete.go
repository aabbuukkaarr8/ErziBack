package product

import "fmt"

const deleteProductQuery = `
DELETE FROM products
WHERE id = $1
`

func (r *Repository) Delete(id int) error {
	res, err := r.store.GetConn().Exec(deleteProductQuery, id)
	if err != nil {
		return fmt.Errorf("repository Delete: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository Delete: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("repository Delete: product not found")
	}

	return nil
}
