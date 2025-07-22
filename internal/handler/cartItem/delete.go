package cartItem

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteCartItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	if err := h.srv.Delete(id); err != nil {
		logrus.WithError(err).Error("[CartItem.Delete]")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot delete item"})
		return
	}

	c.Status(http.StatusOK)
}
