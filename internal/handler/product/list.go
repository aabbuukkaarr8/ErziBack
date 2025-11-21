package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) List(c *gin.Context) {
	products, err := h.srv.List()
	if err != nil {
		logrus.WithError(err).Errorf("[List]Error getting active products")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var response []Model
	for i := range products {
		var p Model
		p.FillFromService(&products[i])
		response = append(response, p)
	}

	c.JSON(http.StatusOK, response)
}
