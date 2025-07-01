package product

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetAll(c *gin.Context) {
	products, err := h.srv.GetAll()
	if err != nil {
		logrus.WithError(err).Errorf("[Get All]Error getting products")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response []Product
	for i := range products {
		var p Product
		p.FillFromService(&products[i])
		response = append(response, p)
	}

	c.JSON(http.StatusOK, response)

}
