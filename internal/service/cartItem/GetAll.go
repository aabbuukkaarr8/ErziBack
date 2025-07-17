package cartItem

type ProductMiniInfo struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}

type CartItemResponse struct {
	ID        int             `json:"id"`
	CartID    int             `json:"cart_id"`
	ProductID int             `json:"product_id"`
	Quantity  int             `json:"quantity"`
	Product   ProductMiniInfo `json:"product"`
}

func (s *Service) GetAll(cartID int) ([]CartItemResponse, error) {
	dbitems, err := s.repo.GetAll(cartID)
	if err != nil {
		return nil, err
	}
	var resp []CartItemResponse
	for _, dbi := range dbitems {
		resp = append(resp, CartItemResponse{
			ID:        dbi.ID,
			CartID:    dbi.CartID,
			ProductID: dbi.ProductID,
			Quantity:  dbi.Quantity,
			Product: ProductMiniInfo{
				Title:    dbi.Product.Title,
				Price:    dbi.Product.Price,
				ImageURL: dbi.Product.ImageURL,
			},
		})
	}
	return resp, nil
}
