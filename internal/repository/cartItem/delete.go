package cartItem

func (r *Repository) Delete(id int) error {
	_, err := r.store.GetConn().Exec(
		`DELETE FROM cart_items WHERE id = $1`, id,
	)
	return err
}
