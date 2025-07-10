package product

import "fmt"

func (r *Repository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	res, err := r.store.GetConn().Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("product not found")
	}
	return nil
}
