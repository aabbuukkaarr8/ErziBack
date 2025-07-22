package cartItem

import "erzi_new/internal/service/cartItem"

func (m *AddCartItemRequest) ToSrv() cartItem.AddCartItemRequest {
	return cartItem.AddCartItemRequest{
		ProductID: m.ProductID,
		UserID:    m.UserID,
	}
}

type AddCartItemRequest struct {
	ProductID int
	UserID    int
}
