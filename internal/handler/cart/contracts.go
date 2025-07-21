package cart

type Service interface {
	GetActive(userID int) (int, error)
}
