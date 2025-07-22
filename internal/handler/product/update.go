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
	return product.UpdateProduct{
		ID:          id,
		Title:       m.Title,
		Description: m.Description,
		Price:       m.Price,
		Quantity:    m.Quantity,
		ImageURL:    m.ImageURL,
		Category:    m.Category,
	}
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

	updatedSrv, err := h.srv.Update(input.ToSrv(id))
	if err != nil {
		logrus.WithError(err).Errorf("[Update ] Error updating product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := Model{}
	result.FillFromService(updatedSrv)

	c.JSON(http.StatusOK, result)
}
