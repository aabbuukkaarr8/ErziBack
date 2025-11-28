package product

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"erzi_new/internal/service/product"
	"erzi_new/pkg/validator"

	"github.com/gin-gonic/gin"
)

func (m *CreateProduct) ToSrv() product.CreateProduct {
	return product.CreateProduct{
		Title:                m.Title,
		Description:          m.Description,
		Price:                m.Price,
		IsActive:             m.IsActive,
		Category:             m.Category,
		BulkDiscountQuantity: m.BulkDiscountQuantity,
		BulkDiscountPrice:    m.BulkDiscountPrice,
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
	createdPSrv, err := h.srv.Create(c.Request.Context(), pSrv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	createdP := Model{}
	createdP.FillFromService(createdPSrv)
	c.JSON(http.StatusCreated, createdP)
}
