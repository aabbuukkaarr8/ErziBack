package product

import "erzi_new/internal/repository/product"

func (m *Product) FillFromDB(dbm *product.Product) {
	m.ID = dbm.ID
	m.Title = dbm.Title
	m.Description = dbm.Description
	m.Price = dbm.Price
	m.ImageURL = dbm.ImageURL
	m.Quantity = dbm.Quantity
	m.Category = dbm.Category
	m.CreatedAt = dbm.CreatedAt
}
