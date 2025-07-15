package cartItem

import cartService "erzi_new/internal/handler/cart"

type Handler struct {
	srv     Service
	cartsrv cartService.Service
}

func NewHandler(srv Service, cartsrv cartService.Service) *Handler {
	return &Handler{
		srv:     srv,
		cartsrv: cartsrv,
	}
}
