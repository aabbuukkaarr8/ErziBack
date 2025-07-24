package cartItem

import (
	"erzi_new/internal/service/cartItem"
	"github.com/google/uuid"
)

func (m *AddCartItemRequest) ToSrv() cartItem.AddCartItemRequest {
	return cartItem.AddCartItemRequest{
		ProductID: m.ProductID,
		UserID:    m.UserID,
	}
}

type AddCartItemRequest struct {
	ProductID int
	UserID    uuid.UUID
}
