package product

import (
	"erzi_new/internal/service/product"
	"erzi_new/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (m *UpdateProduct) ToSrv(id int) product.UpdateProduct {
	upd := product.UpdateProduct{
		ID:          id,
		Title:       m.Title,
		Description: m.Description,
		IsActive:    m.IsActive,
		ImageURL:    m.ImageURL,
		Category:    m.Category,
	}
	if m.Prices != nil {
		prices := make([]product.PriceEntry, len(*m.Prices))
		for i, p := range *m.Prices {
			prices[i] = product.PriceEntry{
				Quantity: p.Quantity,
				Price:    p.Price,
			}
		}
		upd.Prices = &prices
	}
	return upd
}

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Errorf("[strConv.Atoi]writed ID is not a number")
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
		return
	}

	var input UpdateProduct
	err = validator.BindJSON(&input, c.Request)
	if err != nil {
		logrus.WithError(err).Errorf("[Validate] bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSrv, err := h.srv.Update(c.Request.Context(), input.ToSrv(id))
	if err != nil {
		logrus.WithError(err).Errorf("[Update ] Error updating product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := Model{}
	result.FillFromService(updatedSrv)

	c.JSON(http.StatusOK, result)
}
