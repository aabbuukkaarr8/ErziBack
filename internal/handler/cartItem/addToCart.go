package cartItem

import (
	srvcartitem "erzi_new/internal/service/cartItem"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AddCartItemRequest struct {
	ProductID int `json:"product_id" binding:"required"`
}

func (h *Handler) AddCartItem(c *gin.Context) {
	var req AddCartItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithError(err).Errorf("[BindJSON] wrong JSON")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		logrus.Errorf("[c.Get] не найден userID в токене")
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	var intUserID int
	switch v := userID.(type) {
	case float64:
		intUserID = int(v)
	case int:
		intUserID = v
	default:
		c.JSON(401, gin.H{"error": "invalid user id type"})
		return
	}

	cart, err := h.cartsrv.GetCart(intUserID)
	if err != nil {
		logrus.WithError(err).Errorf("[GetCart] cart is not found")
		c.JSON(500, gin.H{"error": "cannot get/create cart"})
		return
	}

	foradd := srvcartitem.AddCartItem{
		ProductID: req.ProductID,
		CartID:    cart,
	}

	added, err := h.srv.Add(foradd)
	if err != nil {
		logrus.WithError(err).Errorf("[Srv.Add] cant add item")
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, added)
}
