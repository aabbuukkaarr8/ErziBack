package cartItem

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteAll(c *gin.Context) {
	idStr := c.Param("cart_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Errorf("[strconv.Atoi] Your id is not int")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.srv.DeleteAll(id)
	if err != nil {
		logrus.WithError(err).Errorf("[srv.DeleteAll] Error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
