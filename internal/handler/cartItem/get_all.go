package cartItem

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetAllCartItems(c *gin.Context) {
	raw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	s, ok := raw.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userID, err := uuid.Parse(s)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	cartID, err := h.cartSrv.GetActive(userID)
	if err != nil {
		logrus.WithError(err).Error("[GetAllCartItems] cannot get cart")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get cart"})
		return
	}
	items, err := h.srv.GetAll(cartID)
	if err != nil {
		logrus.WithError(err).Error("[GetAllCartItems] cannot get items")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}
