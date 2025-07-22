package cartItem

func (r *Repository) DeleteAll(cartID int) error {
	query := "UPDATE carts SET status = 'deleted' WHERE id = $1"
	_, err := r.store.GetConn().Exec(query, cartID)
	return err

}
