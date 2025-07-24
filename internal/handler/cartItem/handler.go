package cartItem

type Handler struct {
	srv     Service
	cartSrv CartService
}

func NewHandler(srv Service, cartSrv CartService) *Handler {
	return &Handler{
		srv:     srv,
		cartSrv: cartSrv,
	}
}
