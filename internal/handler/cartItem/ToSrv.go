package cartItem

import "erzi_new/internal/service/cartItem"

func (m *AddCartItem) ToSrv() cartItem.AddCartItem {
	return cartItem.AddCartItem{
		ProductID: m.ProductID,
		UserID:    m.UserID,
	}
}
