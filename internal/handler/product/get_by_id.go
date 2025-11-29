package product

import (
	"erzi_new/internal/service/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (m *Model) FillFromService(sm *product.Model) {
	m.ID = sm.ID
	m.Title = sm.Title
	m.Description = sm.Description
	m.ImageURL = sm.ImageURL
	m.IsActive = sm.IsActive
	m.Category = sm.Category
	m.CreatedAt = sm.CreatedAt
	m.Prices = make([]PriceEntry, len(sm.Prices))
	for i, p := range sm.Prices {
		m.Prices[i] = PriceEntry{
			Quantity: p.Quantity,
			Price:    p.Price,
		}
	}
}

func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Errorf("[strconv.Atoi] Your ID is not a number")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ps, err := h.srv.GetByID(c.Request.Context(), id)
	if err != nil {
		logrus.WithError(err).Errorf("[Get By ID]Product Not Found")
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	p := Model{}
	p.FillFromService(ps)

	a, err := h.srv.GetAttributes(c.Request.Context(), p.ID)
	if err != nil {
		logrus.WithError(err).Errorf("[Get Attributes]Attributes Not Found")
	}

	for i := range a {
		var m Attribute
		m.FillFromSRV(&a[i])
		p.Attribute = append(p.Attribute, m)
	}
	c.JSON(http.StatusOK, p)

}
