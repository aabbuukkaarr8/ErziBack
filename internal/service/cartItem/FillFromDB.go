package cartItem

import "erzi_new/internal/repository/cartItem"

func (m *CartItemResponse) FillFromDB(dbm *cartItem.CartItemWithProduct) {
	m.ID = dbm.ID
	m.Quantity = dbm.Quantity
	m.CartID = dbm.CartID
	m.ProductID = dbm.ProductID
	m.CreatedAt = dbm.CreatedAt
	m.Product.Title = dbm.Product.Title
	m.Product.Price = dbm.Product.Price
	m.Product.ImageURL = dbm.Product.ImageURL
}
