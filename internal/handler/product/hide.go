package product

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) Hide(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Errorf("[strConv.Atoi]writed ID is not a number")
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный ID"})
		return
	}

	err = h.srv.Hide(c.Request.Context(), id)
	if err != nil {
		logrus.WithError(err).Errorf("[SRV.Switched]something went wrong")
	}
	c.JSON(http.StatusOK, "switched")
}
