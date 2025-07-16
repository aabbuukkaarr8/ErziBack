package product

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Errorf("[strconv.Atoi] writed id is not a number")
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный Id"})
		return
	}
	err = h.srv.Delete(id)
	if err != nil {
		logrus.WithError(err).Errorf("[Delete] Error deleting product")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"message": "товар удален"})
}
