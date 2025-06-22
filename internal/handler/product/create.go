package product

import (
	serviceProduct "erzi_new/internal/service/product"
	"erzi_new/pkg/validator"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateProduct struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=1"`
	Quantity    int     `json:"quantity" validate:"min=0"`
}

func (h *Handler) Create(c *gin.Context) {
	var product serviceProduct.CreateProduct
	err := validator.BindJSON(&product, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdProduct, err := h.srv.Create(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, createdProduct)
}
