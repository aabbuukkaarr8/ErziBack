package cart

type Service interface {
	GetCart(userID int) (int, error)
}
