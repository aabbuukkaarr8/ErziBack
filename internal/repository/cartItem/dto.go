package cartItem

import "time"

type Model struct {
	ID        int
	CartID    int
	ProductID int
	Quantity  int
	CreatedAt time.Time
}
type ProductMiniInfo struct {
	Title    string
	Price    float64
	ImageURL string
}
type ModelWithProduct struct {
	ID        int
	CartID    int
	ProductID int
	Quantity  int
	CreatedAt time.Time
	Product   ProductMiniInfo
}
