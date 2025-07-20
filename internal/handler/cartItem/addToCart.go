package cartItem

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) AddCartItem(c *gin.Context) {
	idStr := c.Param("product_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		logrus.WithError(err).Errorf("[strconv.Atoi] Your id is not int")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		logrus.Errorf("[c.Get] не найден userID в токене")
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	var intUserID int
	switch v := userID.(type) {
	case float64:
		intUserID = int(v)
	case int:
		intUserID = v
	default:
		c.JSON(401, gin.H{"[UserID]": "invalid user id type"})
		return
	}

	cart, err := h.cartsrv.GetCart(intUserID)
	if err != nil {
		logrus.WithError(err).Errorf("[GetCart] cart is not found")
		c.JSON(500, gin.H{"error": "cannot get/create cart"})
		return
	}
	i := AddCartItem{
		ProductID: id,
		CartID:    cart,
	}
	iSrv := i.ToSrv()
	_, err = h.srv.Add(iSrv)
	if err != nil {
		logrus.WithError(err).Errorf("[Srv.Add] cant add item")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, nil)
}
