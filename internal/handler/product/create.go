package product

import (
	"github.com/sirupsen/logrus"
	"net/http"

	"erzi_new/internal/service/product"
	"erzi_new/pkg/validator"
	"github.com/gin-gonic/gin"
)

type CreateProduct struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,min=1"`
	Quantity    int     `json:"quantity" validate:"min=0"`
	Category    string  `json:"category" validate:"required,oneof=honey-jam meltwater mineral-water equipment"`
}

func (m *CreateProduct) ToSrv() product.CreateProduct {
	return product.CreateProduct{
		Title:       m.Title,
		Description: m.Description,
		Price:       m.Price,
		Quantity:    m.Quantity,
		Category:    m.Category,
	}
}

func (h *Handler) Create(c *gin.Context) {
	var p CreateProduct
	err := validator.BindJSON(&p, c.Request)
	if err != nil {
		logrus.WithError(err).Warn("Product.Create: invalid request JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pSrv := p.ToSrv()
	createdPSrv, err := h.srv.Create(pSrv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdP := Product{}
	createdP.FillFromService(createdPSrv)
	c.JSON(http.StatusCreated, createdP)
}
