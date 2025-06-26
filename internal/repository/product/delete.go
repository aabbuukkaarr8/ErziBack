package product

func (r *Repository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.store.GetConn().Exec(query, id)
	return err
}
