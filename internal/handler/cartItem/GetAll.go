package cartItem

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) GetAllCartItems(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		logrus.Error("[GetAllCartItems] userID not found in JWT")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	var uid int
	switch v := userID.(type) {
	case float64:
		uid = int(v)
	case int:
		uid = v
	default:
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user id type"})
		return
	}
	cartID, err := h.cartsrv.GetActive(uid)
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
